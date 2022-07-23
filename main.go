package main

import (
	"log"
	"net"

	"github.com/go-playground/validator/v10"
	"github.com/vincen320/user-service-grpc/app"
	"github.com/vincen320/user-service-grpc/controller"
	pb "github.com/vincen320/user-service-grpc/proto"
	"github.com/vincen320/user-service-grpc/repository"
	"github.com/vincen320/user-service-grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	DB, err := app.ConnectMongo()
	if err != nil {
		panic(err)
	}
	validator := validator.New()
	app.RegisterValidatorTagName(validator)

	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepository, validator)

	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// Register reflection service on gRPC server.
	// https://stackoverflow.com/questions/41424630/why-do-we-need-to-register-reflection-service-on-grpc-server
	// https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	// YT : [Backend #41] How to run a golang gRPC server and call its API => Minutes 4:44
	reflection.Register(grpcServer)

	userGrpcServer := controller.NewUserGrpcServer(userService)
	pb.RegisterUserServer(grpcServer, userGrpcServer)

	log.Println("user grpc server start at", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot connect grpc server")
	}
}
