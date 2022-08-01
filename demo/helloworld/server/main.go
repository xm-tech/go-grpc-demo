package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/xm-tech/go-grpc-demo/api/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

// Sends a greeting
func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Println("server.SayHello, req=", req.GetName())
	msg := fmt.Sprintf("hello %v", req.GetName())
	return &helloworld.HelloReply{Message: msg}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at : %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
