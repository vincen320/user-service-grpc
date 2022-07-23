package controller

import (
	"context"

	"github.com/vincen320/user-service-grpc/helper"
	pb "github.com/vincen320/user-service-grpc/proto"
	"github.com/vincen320/user-service-grpc/service"
)

type UserGrpcServer struct {
	Service service.UserService
	pb.UnimplementedUserServer
}

func NewUserGrpcServer(service service.UserService) pb.UserServer {
	return &UserGrpcServer{
		Service: service,
	}
}

func (us *UserGrpcServer) CreateUser(ctx context.Context, userCreateRequest *pb.UserCreateRequest) (*pb.UserCreateResponse, error) {
	request := helper.ProtoCreateUserToWebCreateUser(userCreateRequest)

	response, err := us.Service.CreateUser(ctx, request)
	if err != nil {
		return nil, helper.ReturnProtoError(err)
	}

	return helper.UserResponseToProtoCreateUserResponse(response), nil
}

func (us *UserGrpcServer) GetUser(ctx context.Context, getUserReqeust *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	request := getUserReqeust.GetSearch()

	response, err := us.Service.FindUser(ctx, request)
	if err != nil {
		return nil, helper.ReturnProtoError(err)
	}

	return helper.UserResponseToProtoGetUserResponse(response), nil
}
