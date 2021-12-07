package services

import (
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, dto CreateUserDto) (userDto *UserDto, err error)
	FindUserByID(ctx context.Context, userId string) (userDto *UserDto, err error)
	FindAllUser(ctx context.Context) (usersDto []UserDto, err error)
}

type userService struct {
}

func NewUserService() (UserService, error) {

	return &userService{}, nil
}

func (service *userService) CreateUser(ctx context.Context, dto CreateUserDto) (userDto *UserDto, err error) {
	return &UserDto{
		UserID:   "1",
		UserName: dto.Email,
		Email:    dto.Email,
		Password: dto.Password,
	}, nil
}

func (service *userService) FindAllUser(ctx context.Context) (usersDto []UserDto, err error) {
	return users, nil
}

func (service *userService) FindUserByID(ctx context.Context, userId string) (userDto *UserDto, err error) {
	return &UserDto{
		UserID:   "1",
		UserName: "ann",
		Email:    "ann@email.com",
		Password: "1234",
	}, nil
}

type CreateUserDto struct {
	Password string `json:"password" form:"password" binding:"required,password"`
	Email    string `json:"email" form:"email" binding:"required,email"`
}

type UserDto struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserServiceError struct {
	ServiceError
}

func NewUserServiceError(e ServiceEvent) error {
	return &UserServiceError{ServiceError{ServiceName: "wrapUserService", Code: e.GetEvent().Code, Msg: e.GetEvent().Msg, Err: nil}}
}

var users = []UserDto{
	{
		UserID:   "1",
		UserName: "ann",
		Email:    "ann@email.com",
		Password: "1234",
	},
	{
		UserID:   "2",
		UserName: "wei",
		Email:    "wei@email.com",
		Password: "1234",
	},
	{
		UserID:   "3",
		UserName: "woo",
		Email:    "woo@email.com",
		Password: "1234",
	},
	{
		UserID:   "4",
		UserName: "dad",
		Email:    "dad@email.com",
		Password: "1234",
	},
	{
		UserID:   "5",
		UserName: "mon",
		Email:    "mon@email.com",
		Password: "1234",
	},
}
