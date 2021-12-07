package grpc

import (
	"fmt"
	pb "go-api-protocols/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var (
	routerInstance *Router
)

type Router struct {
	server *grpc.Server
}

func GetRouter() (instance *Router, err error) {
	if routerInstance == nil {
		routerInstance, err = newRouter()
		if err != nil {
			return nil, err
		}
	}
	return routerInstance, nil
}

func newRouter() (router *Router, err error) {
	server := grpc.NewServer()
	user, err := NewUserService()
	if err != nil {
		return nil, err
	}
	pb.RegisterUserServiceServer(server, user)
	reflection.Register(server)
	return &Router{
		server,
	}, nil
}

func (r *Router) Run(address string) error {
	fmt.Printf("[g-RPC-Debug] Listening and serving HTTP on %s\n", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	err = r.server.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
