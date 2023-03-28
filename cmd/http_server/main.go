package main

import (
	"gin-practice/apps/hello/access"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"luwu-protobuf/hello_protobuf"
	"net"
)

func main() {
	go startGrpc()
	engine := gin.Default()
	gin.Recovery()
	_ = engine.SetTrustedProxies(nil)
	if e := InitApp(engine); e != nil { // register urls of all app
		log.Fatalf("Application register failed, err = %v", e)
		return
	}
	if e := engine.Run("127.0.0.1:8080"); e != nil {
		log.Fatalf("Gin engine run failed, err = %v", e)
		return
	}
}

func startGrpc() {
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
