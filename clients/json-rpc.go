package clients

import (
	"bytes"
	"fmt"
	"github.com/gorilla/rpc/json"
	jsonrpcServices "go-api-protocols/adapter/jsonrpc/services"
	"net/http"
)

type JsonRpcClient struct {
	*http.Client
}

func NewJsonRpcClient() (Client, error) {
	return JsonRpcClient{
		&http.Client{},
	}, nil
}

func (client JsonRpcClient) Execute() (err error) {
	args := &jsonrpcServices.FindUserArgs{
		Id: "0",
	}
	message, err := json.EncodeClientRequest("UserService.FindUser", args)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", "http://localhost:8083/rpc", bytes.NewBuffer(message))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var r jsonrpcServices.UserReply
	err = json.DecodeClientResponse(resp.Body, &r)
	if err != nil {
		return
	}
	fmt.Println("json-rpc response:", r)
	return nil
}
