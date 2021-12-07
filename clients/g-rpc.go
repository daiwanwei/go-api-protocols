package clients

import (
	"context"
	"fmt"
	pb "go-api-protocols/protos"
	"google.golang.org/grpc"
	"time"
)

type GrpcClient struct {
}

func NewGrpcClient() (Client, error) {
	return GrpcClient{}, nil
}

func (client GrpcClient) Execute() (err error) {
	conn, err := grpc.Dial("localhost:8084", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := c.FindUser(ctx, &pb.FindUserRequest{Id: "0"})
	if err != nil {
		return err
	}
	fmt.Println("g-rpc response:", res)
	return nil
}
