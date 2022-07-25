package repository

import (
	"context"
	"errors"

	"github.com/vincen320/user-service-grpc/exception"
	"github.com/vincen320/user-service-grpc/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	DB *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (us *UserRepositoryImpl) Save(ctx context.Context, User domain.User) (domain.User, error) {
	result, err := us.DB.Collection("user").InsertOne(ctx, User)
	if err != nil {
		return User, err
	}
	userId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return User, errors.New("unable convert ID") // 500 Internal Server Error
	}
	User.Id = userId.Hex() //Gunakan .Hex() dibanding .String()
	return User, nil
}

func (us *UserRepositoryImpl) Find(ctx context.Context, User domain.User) (domain.User, error) {
	filter := filterBy(&User)
	result := us.DB.Collection("user").FindOne(ctx, filter)

	var user domain.User
	err := result.Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			//untuk pesan error saja biar enak
			var with string
			if filter["username"] != nil { //pake nil karena map dari primitive.M adalah map[string]interface{} , valuenya interface{} jadi pengecekkannya adalah nil
				with = "username " + filter["username"].(string)
			} else if filter["email"] != nil {
				with = "email " + filter["email"].(string)
			} else {
				//berarti pakai username dan email, harus ekstrak "$or" nya dulu (dalam kasus ini hanya createUser yang menggunakan ini)
				filter2, _ := filter["$or"].([]bson.M)
				with = "username " + (filter2[0]["username"]).(string) + " and email " + (filter2[1]["email"]).(string)
			}
			//end
			return user, exception.NewNotFoundError("user with " + with + " not found") //404 not found, return error trs smpe service ttp return error
		}
		return user, err // 500 internal server error || unknown error
	}
	return user, nil
}

func filterBy(User *domain.User) bson.M {
	if User.Email != "" && User.Username != "" {
		return bson.M{
			"$or": []bson.M{
				{
					"username": User.Username,
				}, {
					"email": User.Email,
				},
			},
		}
	}

	if User.Email != "" {
		return bson.M{
			"email": User.Email,
		}
	}

	return bson.M{
		"username": User.Username,
	}
}
