package service

import (
	"context"

	"github.com/LandvibeDev/gofka-codelab-sample/errors"
	"github.com/LandvibeDev/gofka-codelab-sample/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceInterface interface {
	GetByID(*context.Context, string) (*model.User, error)
	GetByName(*context.Context, string) (*model.User, error)
	GetByEmail(*context.Context, string) (*model.User, error)
	Create(*context.Context, *model.User) (*model.User, error)
	Update(*context.Context, *model.User) (*model.User, error)
	Delete(*context.Context, string) error
}

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(client *mongo.Client) *UserService {
	return &UserService{
		collection: client.Database("gofka").Collection("users"),
	}
}

func (us *UserService) GetByID(ctx *context.Context, id string) (*model.User, error) {
	var user model.User
	if err := us.collection.FindOne(*ctx, bson.D{{"id", id}}).Decode(&user); err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func (us *UserService) GetByName(ctx *context.Context, name string) (*model.User, error) {
	var user model.User
	if err := us.collection.FindOne(*ctx, bson.D{{"name", name}}).Decode(&user); err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func (us *UserService) GetByEmail(ctx *context.Context, email string) (*model.User, error) {
	var user model.User
	if err := us.collection.FindOne(*ctx, bson.D{{"email", email}}).Decode(&user); err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func (us *UserService) Create(ctx *context.Context, u *model.User) (*model.User, error) {
	_, err := us.collection.InsertOne(*ctx, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (us *UserService) Update(ctx *context.Context, u *model.User) (*model.User, error) {
	result, err := us.collection.ReplaceOne(*ctx, bson.D{{"id", u.ID}}, u)
	if err != nil || result == nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.NotFound{ID: u.ID}
	} else {
		return u, nil
	}
}

func (us *UserService) Delete(ctx *context.Context, id string) error {
	result, err := us.collection.DeleteOne(*ctx, bson.D{{"id", id}})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.NotFound{ID: id}
	} else {
		return nil
	}
}
