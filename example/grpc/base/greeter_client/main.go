package main

/**
* @Time       : 2020/11/22 7:11 下午
* @Author     : xumamba
* @Description: Greeter client
 */

import (
	"bufio"
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "go-life/example/grpc/base"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to creates client connection: %v", err)
	}
	defer clientConn.Close()
	client := pb.NewGreeterClient(clientConn)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
		defer cancelFunc()
		text := scanner.Text()
		reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: text})
		if err != nil {
			log.Fatalf("failed to say hello: %v", err)
		}
		log.Printf("Greetings from success: %v", reply.Message)
	}

}
