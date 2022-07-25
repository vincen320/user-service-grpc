package app

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/vincen320/user-service-grpc/controller"
	"github.com/vincen320/user-service-grpc/proto"
	pb "github.com/vincen320/user-service-grpc/proto"
	"github.com/vincen320/user-service-grpc/repository"
	"github.com/vincen320/user-service-grpc/service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	tls            = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile       = flag.String("cert_file", "", "The TLS cert file")
	keyFile        = flag.String("key_file", "", "The TLS key file")
	DB             *mongo.Database
	validators     *validator.Validate
	userRepository repository.UserRepository
	userService    service.UserService
	userGrpcServer proto.UserServer
)

func initDB() {
	if DB == nil {
		DB = ConnectMongo()
	}
}

func initValidator() {
	if validators == nil {
		validators = validator.New()
		RegisterValidatorTagName(validators)
	}
}

func initUserRepository() {
	if userRepository == nil {
		userRepository = repository.NewUserRepository(DB)
	}
}

func initUserService() {
	if userService == nil {
		userService = service.NewUserService(userRepository, validators)
	}
}

func initUserGrpcServer() {
	if userGrpcServer == nil {
		userGrpcServer = controller.NewUserGrpcServer(userService)
	}
}

func initDependencies() {
	initDB()
	initValidator()
	initUserRepository()
	initUserService()
	initUserGrpcServer() //struct gRPC Server
}

func StartGrpcServer() {
	flag.Parse()
	initDependencies()

	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//UPDATE : Connection with TLS
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = "ssl/server.crt"
		}
		if *keyFile == "" {
			*keyFile = "ssl/server.pem"
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	grpcServer := grpc.NewServer(opts...)
	// Register reflection service on gRPC server.
	// https://stackoverflow.com/questions/41424630/why-do-we-need-to-register-reflection-service-on-grpc-server
	// https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	// YT : [Backend #41] How to run a golang gRPC server and call its API => Minutes 4:44
	reflection.Register(grpcServer)

	pb.RegisterUserServer(grpcServer, userGrpcServer)

	//gracefully shutdown
	go func() {
		log.Println("user grpc server start at", listener.Addr().String(), "with SSL/TLS =", *tls)
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal("cannot connect grpc server")
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Println("Stopping gRPCServer")
	grpcServer.Stop()
	log.Println("Closing the Listener")
	listener.Close()
}

//HOW : https://github.com/grpc-ecosystem/grpc-gateway/
//https://grpc-ecosystem.github.io/grpc-gateway/
func StartGrpcGatewayServer() {
	flag.Parse()
	initDependencies()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	}) //parameter untuk ganti type camelCase field json jadi snake_case

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible  dan pastikan runtimenya yang diimport v2
	mux := runtime.NewServeMux(jsonOption)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterUserHandlerFromEndpoint(ctx, mux, "localhost:9090", opts) //harus dijalankan server grpcNya, jadi 2x jalan

	if err != nil {
		log.Fatal("error registering handler", err)
	}

	var server http.Server
	//gracefully shutdown
	go func() {
		server = http.Server{
			Addr:    ":8080",
			Handler: mux,
		}
		log.Println("grpc gateway server start at", server.Addr)
		server.ListenAndServe()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Println("closing server")
	server.Close()
}
