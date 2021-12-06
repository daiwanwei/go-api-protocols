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

func (service *UserService) FindAllUser(r *http.Request, args *FindAllUserArgs, reply *UserList) error {
	users, err := service.User.FindAllUser(context.Background())
	if err != nil {
		return err
	}
	*reply = users
	return nil
}

type FindAllUserArgs struct {
}

type UserList []services.UserDto
