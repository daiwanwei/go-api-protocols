package grpc

import (
	"context"
	"github.com/jinzhu/copier"
	"go-api-protocols/business/services"
	pb "go-api-protocols/protos"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	wrapped services.UserService
}

func NewUserService() (pb.UserServiceServer, error) {
	service, err := services.GetService()
	if err != nil {
		return nil, err
	}
	return &UserService{
		UnimplementedUserServiceServer: pb.UnimplementedUserServiceServer{},
		wrapped:                        service.User,
	}, nil
}

func (service *UserService) FindAllUser(ctx context.Context, req *pb.FindAllUserRequest) (res *pb.UserListResponse, err error) {
	users, err := service.wrapped.FindAllUser(context.Background())
	if err != nil {
		return nil, err
	}
	res = &pb.UserListResponse{}
	err = copier.Copy(&res.Users, &users)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *UserService) FindUser(ctx context.Context, req *pb.FindUserRequest) (res *pb.UserResponse, err error) {
	user, err := service.wrapped.FindUserByID(context.Background(), req.Id)
	if err != nil {
		return nil, err
	}
	res = &pb.UserResponse{
		User: &pb.User{},
	}
	err = copier.Copy(res.User, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
