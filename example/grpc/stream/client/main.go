package main

import (
	"log"

	"google.golang.org/grpc"

	pb "go-life/example/grpc/base"
)

/**
* @Time       : 2020/11/22 9:08 下午
* @Author     : xumamba
* @Description: stream example client
 */

func main() {
	clientConn, err := grpc.Dial("127.0.0.1:50002", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatalf("failed to dial: %v", err)
	}
	defer clientConn.Close()
	client := pb.NewGreeterClient(clientConn)

}