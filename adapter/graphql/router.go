package graphql

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"go-api-protocols/adapter/graphql/graph"
	"go-api-protocols/adapter/graphql/graph/generated"
	"go-api-protocols/adapter/graphql/middlewares"
	"go-api-protocols/business/services"
	"net/http"
)

var (
	routerInstance *wrapRouter
)

type wrapRouter struct {
	router chi.Router
}

func GetRouter() (instance *wrapRouter, err error) {
	if routerInstance == nil {
		instance, err = newRouter()
		if err != nil {
			return nil, err
		}
		routerInstance = instance
	}
	return routerInstance, nil
}

func newRouter() (*wrapRouter, error) {
	r := chi.NewRouter()
	middleware, err := middlewares.GetMiddleware()
	if err != nil {
		return nil, err
	}
	r.Use(middleware.Auth.Authenticate())
	service, err := services.GetService()
	if err != nil {
		return nil, err
	}
	handlers := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					User: service.User,
					Auth: service.Auth,
				},
			},
		),
	)
	r.Handle("/graphql", playground.Handler("GraphQL playground", "/graphql/query"))
	r.Handle("/graphql/query", handlers)
	return &wrapRouter{r}, nil
}

func (wrap *wrapRouter) Run(address string) error {
	fmt.Printf("[gql-Debug] Listening and serving HTTP on %s\n", address)
	err := http.ListenAndServe(address, wrap.router)
	if err != nil {
		return err
	}
	return nil
}
