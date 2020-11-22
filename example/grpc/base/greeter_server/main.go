package main

/**
* @Time       : 2020/11/22 7:11 下午
* @Author     : xumamba
* @Description: Greeter server
 */

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "go-life/example/grpc/base"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Reveived greeter: %v", request.Name)
	return &pb.HelloReply{Message: "Hello " + request.GetName()}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
