package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"go-api-protocols/adapter/graphql/graph"
	"go-api-protocols/adapter/graphql/graph/generated"
	"go-api-protocols/business/services"
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
