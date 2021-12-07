package grpc

import (
	"context"
	"github.com/jinzhu/copier"
	"go-api-protocols/business/services"
	pb "go-api-protocols/protos"
)

type wrapUserService struct {
	services.UserService
}

func NewUserService() (pb.UserServiceServer, error) {
	service, err := services.GetService()
	if err != nil {
		return nil, err
	}
	return &wrapUserService{service.User}, nil
}

func (wrap *wrapUserService) FindAllUser(ctx context.Context, req *pb.FindAllUserRequest) (res *pb.UserListResponse, err error) {
	users, err := wrap.UserService.FindAllUser(context.Background())
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

func (wrap *wrapUserService) FindUser(ctx context.Context, req *pb.FindUserRequest) (res *pb.UserResponse, err error) {
	user, err := wrap.UserService.FindUserByID(context.Background(), req.Id)
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
