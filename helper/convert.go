package helper

import (
	"time"

	"github.com/vincen320/user-service-grpc/model/domain"
	"github.com/vincen320/user-service-grpc/model/web"
	pb "github.com/vincen320/user-service-grpc/proto"
)

func CreateUserRequestToUser(createUserRequest web.CreateUserRequest) domain.User {
	return domain.User{
		Name:      createUserRequest.Name,
		Username:  createUserRequest.Username,
		Password:  createUserRequest.Password,
		Email:     createUserRequest.Email,
		CreatedAt: time.Now().UTC().Unix(),
	}
}

func UserToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func ProtoCreateUserToWebCreateUser(fromProto *pb.UserCreateRequest) web.CreateUserRequest {
	return web.CreateUserRequest{
		Name:     fromProto.GetName(),
		Username: fromProto.GetUsername(),
		Password: fromProto.GetPassword(),
		Email:    fromProto.GetEmail(),
	}
}

func UserResponseToProtoCreateUserResponse(response web.UserResponse) *pb.UserCreateResponse {
	return &pb.UserCreateResponse{
		Id:        response.Id,
		Name:      response.Name,
		Username:  response.Username,
		Email:     response.Email,
		CreatedAt: response.CreatedAt,
		Message:   "Success Create User",
	}
}

func UserResponseToProtoGetUserResponse(response web.UserResponse) *pb.GetUserResponse {
	return &pb.GetUserResponse{
		Id:        response.Id,
		Name:      response.Name,
		Username:  response.Username,
		Email:     response.Email,
		CreatedAt: response.CreatedAt,
		Message:   "Success Find User",
	}
}
