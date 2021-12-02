package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"go-api-protocols/adapter/graphql/graph/generated"
	"go-api-protocols/adapter/graphql/graph/model"
	"go-api-protocols/adapter/graphql/middlewares"
	"go-api-protocols/business/services"

	"github.com/jinzhu/copier"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
	user, err := r.User.CreateUser(context.Background(), services.CreateUserDto{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	userModel := &model.User{}
	err = copier.Copy(&userModel, user)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	passport, err := r.Auth.Login(context.Background(), services.LoginDto{
		UserName: input.UserName,
		Password: input.Password,
	})
	if err != nil {
		return "", err
	}
	return passport.Token, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	auth := middlewares.GetAuthenticationForContext(ctx)
	if auth == nil {
		return nil, errors.New("don't have authority")
	} else {
		fmt.Println(auth.GetName())
	}
	users, err := r.User.FindAllUser(context.Background())
	if err != nil {
		return nil, err
	}
	var usersModel []*model.User
	err = copier.Copy(&usersModel, users)
	if err != nil {
		return nil, err
	}
	return usersModel, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
