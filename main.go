package main

import (
	"fmt"
	"go-api-protocols/adapter/graphql"
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

	gqlServer, err := graphql.NewServer()
	go func() {
		defer wg.Done()
		err = gqlServer.Run(":8082")
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	gqlRouter, err := graphql.GetRouter()
	go func() {
		defer wg.Done()
		err = gqlRouter.Run(":8083")
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	wg.Wait()
}
