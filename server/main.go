package main

import (
	"fmt"
	"net"

	pb "github.com/Aytaditya/grpc-learning/proto"
	"google.golang.org/grpc"
)

const port = ":8080"

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	ls, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	fmt.Println("gRPC server is running on port", ls.Addr())
	err = grpcServer.Serve(ls)
	if err != nil {
		panic(err)
	}
}
