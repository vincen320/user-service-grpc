package repository

import (
	"context"

	"github.com/vincen320/user-service-grpc/model/domain"
)

type UserRepository interface {
	Save(context.Context, domain.User) (domain.User, error)
	Find(context.Context, domain.User) (domain.User, error)
}
