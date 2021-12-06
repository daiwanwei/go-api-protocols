package main

import (
	"fmt"
	"go-api-protocols/adapter/graphql"
	"go-api-protocols/adapter/jsonrpc/routers"
	restRouters "go-api-protocols/adapter/rest/routers"
	_ "go-api-protocols/docs"
	"sync"
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
	wg.Add(3)
	restRouter, err := restRouters.GetRouter()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	go func() {
		defer wg.Done()
		err = restRouter.Run(":8080")
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
		err = gqlRouter.Run(":8081")
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	jsonRpcRouter, err := routers.GetRouter()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	go func() {
		defer wg.Done()
		err = jsonRpcRouter.Run(":8082")
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	wg.Wait()
}
