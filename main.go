package main

import (
	"github.com/vincen320/user-service-grpc/app"
)

func main() {
	go app.StartGrpcServer() //harus jalan untuk grpcGatewayServer
	app.StartGrpcGatewayServer()
}
