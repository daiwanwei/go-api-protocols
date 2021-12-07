package main

import (
	"fmt"
	"go-api-protocols/adapter/graphql"
	"go-api-protocols/adapter/grpc"
	"go-api-protocols/adapter/jsonrpc"
	restRouters "go-api-protocols/adapter/rest/routers"
	"go-api-protocols/clients"
	_ "go-api-protocols/docs"
	"sync"
	"time"
)

// @title Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email wadejet.work@gmail.com

// @license.name MIT
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @BasePath /
func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	restRouter, err := restRouters.GetRouter()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	go func() {
		defer wg.Done()
		err = restRouter.Run(":8081")
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	gqlRouter, err := graphql.GetRouter()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	go func() {
		defer wg.Done()
		err = gqlRouter.Run(":8082")
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	jsonRpcRouter, err := jsonrpc.GetRouter()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	go func() {
		defer wg.Done()
		err = jsonRpcRouter.Run(":8083")
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	grpcRouter, err := grpc.GetRouter()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	go func() {
		defer wg.Done()
		err = grpcRouter.Run(":8084")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("grpc router close")
	}()
	time.Sleep(time.Second * 3)
	jsonRpcClient, err := clients.NewJsonRpcClient()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = jsonRpcClient.Execute()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	grpcClient, err := clients.NewGrpcClient()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = grpcClient.Execute()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	wg.Wait()
}
