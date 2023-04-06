package main

import (
	"github.com/memorate/luwu-protobuf/hello_protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"luwu/apps/hello/access"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello_protobuf.RegisterHelloServiceServer(s, access.NewHelloGrpcApp())

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
