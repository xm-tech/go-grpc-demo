package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/xm-tech/go-grpc-demo/api/helloworld"
	"google.golang.org/grpc"
)

func main() {
	// set up a conn to the server
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建1客户端连接的代理
	agent := helloworld.NewGreeterClient(conn)

	// 解析参数
	txt := "world"
	if len(os.Args) > 1 {
		txt = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := agent.SayHello(ctx, &helloworld.HelloRequest{Name: txt})
	if err != nil {
		panic(err)
	}

	log.Println("resp: ", r.GetMessage())
}
