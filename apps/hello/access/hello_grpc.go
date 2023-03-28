package access

import (
	"context"
	"fmt"
	"luwu-protobuf/hello_protobuf"
)

type HelloGrpcApp struct {
}

func NewHelloGrpcApp() *HelloGrpcApp {
	return &HelloGrpcApp{}
}

func (h *HelloGrpcApp) SayHello(ctx context.Context, request *hello_protobuf.HelloRequest) (*hello_protobuf.HelloResponse, error) {
	message := fmt.Sprintf("Hi, %s, this is LuWu ^_^", request.Name)
	return &hello_protobuf.HelloResponse{
		Message: message,
	}, nil
}
