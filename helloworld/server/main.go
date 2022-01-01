package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/xm-tech/go-grpc-demo/helloworld/helloworld"
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
	log.Println("SayHello, req.Name=", req.Name)
	msg := fmt.Sprintf("hello %v", req.Name)
	return &helloworld.HelloReply{Message: msg}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	msg := fmt.Sprintf("hello2 %v", in.Name)
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
