package jsonrpc

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"go-api-protocols/adapter/jsonrpc/services"
	"net/http"
)

var (
	routerInstance *router
)

type router struct {
	wrappedRouter *mux.Router
}

func GetRouter() (instance *router, err error) {
	if routerInstance == nil {
		instance, err = newRouter()
		if err != nil {
			return nil, err
		}
		routerInstance = instance
	}
	return routerInstance, nil
}

func newRouter() (*router, error) {
	r := mux.NewRouter()
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	service, err := services.GetService()
	if err != nil {
		return nil, err
	}
	err = s.RegisterService(service.User, "")
	if err != nil {
		return nil, err
	}
	r.Handle("/rpc", s)
	return &router{r}, nil
}

func (r *router) Run(address string) error {
	fmt.Printf("[json-RPC-Debug] Listening and serving HTTP on %s\n", address)
	err := http.ListenAndServe(address, r.wrappedRouter)
	if err != nil {
		return err
	}
	return nil
}
