package services

import (
	"context"
	"go-api-protocols/business/services"
	"net/http"
)

type UserService struct {
	User services.UserService
}

func NewUserService() (*UserService, error) {
	service, err := services.GetService()
	if err != nil {
		return nil, err
	}
	return &UserService{
		service.User,
	}, nil
}

func (service *UserService) FindAllUser(r *http.Request, args *FindAllUserArgs, reply *UserListReply) error {
	users, err := service.User.FindAllUser(context.Background())
	if err != nil {
		return err
	}
	*reply = users
	return nil
}

func (service *UserService) FindUser(r *http.Request, args *FindUserArgs, reply *UserReply) error {
	user, err := service.User.FindUserByID(context.Background(), args.Id)
	if err != nil {
		return err
	}
	*reply = UserReply(*user)
	return nil
}

type FindAllUserArgs struct {
}

type FindUserArgs struct {
	Id string `json:"id"`
}

type UserListReply []services.UserDto
type UserReply services.UserDto
