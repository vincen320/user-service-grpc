package service

import (
	"context"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/vincen320/user-service-grpc/exception"
	"github.com/vincen320/user-service-grpc/helper"
	"github.com/vincen320/user-service-grpc/model/domain"
	"github.com/vincen320/user-service-grpc/model/web"
	"github.com/vincen320/user-service-grpc/repository"
)

var UserNil domain.User

type UserServiceImpl struct {
	Repository repository.UserRepository
	Validator  *validator.Validate
}

func NewUserService(repo repository.UserRepository, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		Repository: repo,
		Validator:  validator,
	}
}

func (us *UserServiceImpl) CreateUser(ctx context.Context, createUserRequest web.CreateUserRequest) (web.UserResponse, error) {
	err := us.Validator.Struct(createUserRequest)
	if err != nil {
		return web.UserResponse{}, err
	}

	//Find User, isExist akan return object domain.User, yangmana jika propertiesnya ada isi, berarti user ada, kalau tidak ada isi berarti tidak ada
	_, err = us.Repository.Find(ctx, domain.User{
		Email:    createUserRequest.Email,
		Username: createUserRequest.Username,
	})

	//CEK ERRORNYA, KALAU NIL BERARTI AMAN-AMAN AJA == KETEMU USER, KETEMU USER == TIDAK VALID, TIDAK VALID == USER ALREADY EXISTS
	if err == nil {
		return web.UserResponse{}, exception.NewConflictError("user already exists")
	}

	if err != nil {
		var notfound *exception.NotFoundError
		//kalau errornya bukan notfound, berarti error lain
		if !errors.As(err, &notfound) {
			return web.UserResponse{}, err
		}
	}

	//bcryptpassword
	createUserRequest.Password, err = helper.BcryptPassword(createUserRequest.Password)
	if err != nil {
		return web.UserResponse{}, err
	}

	//insert to db
	response, err := us.Repository.Save(ctx, helper.CreateUserRequestToUser(createUserRequest))
	if err != nil {
		return web.UserResponse{}, err
	}

	return helper.UserToUserResponse(response), nil
}

func (us *UserServiceImpl) FindUser(ctx context.Context, search string) (web.UserResponse, error) {
	user := domain.User{}
	if isEmail(search) {
		user.Email = search
	} else {
		user.Username = search
	}

	response, err := us.Repository.Find(ctx, user)
	if err != nil {
		return web.UserResponse{}, err
	}

	return helper.UserToUserResponse(response), nil
}

func isEmail(s string) bool {
	split := strings.Split(s, "@")
	if len(split) > 1 {
		return true
	}
	return false
}
