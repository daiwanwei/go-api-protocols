package graphql

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"go-api-protocols/adapter/graphql/graph"
	"go-api-protocols/adapter/graphql/graph/generated"
	"go-api-protocols/business/services"
	"net/http"
)

type Server struct {
	handlers *handler.Server
}

func NewServer() (server *Server, err error) {
	service, err := services.GetService()
	if err != nil {
		return nil, err
	}
	handlers := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					User: service.User,
				},
			},
		),
	)
	return &Server{
		handlers: handlers,
	}, nil
}

func (s *Server) Run(address string) error {
	http.Handle("/graphql", playground.Handler("GraphQL playground", "/graphql/query"))
	http.Handle("/graphql/query", s.handlers)
	fmt.Printf("[gql-debug]:Listening and serving HTTP on %s \n", address)
	fmt.Printf("[gql-debug]:graph UI endpoint:%s \n", "/graphql")
	fmt.Printf("[gql-debug]:query endpoint:%s \n", "/graphql/query")
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return nil
}
