package service

import (
	"context"

	"github.com/vincen320/user-service-grpc/model/web"
)

type UserService interface {
	CreateUser(context.Context, web.CreateUserRequest) (web.UserResponse, error)
	FindUser(context.Context, string) (web.UserResponse, error)
}
