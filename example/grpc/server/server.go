package main

/**
 * @DateTime   : 2020/11/18
 * @Author     : xumamba
 * @Description:
 **/

import (
	"log"
	"net"
	"os"
	"strings"

	eg "go-life/example/grpc"

	"github.com/gogo/protobuf/proto"
)

const port = ":50051"

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalf("failed to accept:%v", err)
		}
		go simpleService(conn)
	}
}

func simpleService(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096, 4096)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		request := &eg.HelloRequest{}
		reqData := buf[:count]

		err = proto.Unmarshal(reqData, request)
		if err != nil {
			panic(err)
		}

		log.Printf("receive args: %v from: %v\n", request.Greeting, conn.RemoteAddr())
		if request.Greeting == "EOF" {
			os.Exit(1)
		}

		response := &eg.HelloResponse{
			Reply: strings.ToUpper(request.Greeting),
		}
		bytes, err := proto.Marshal(response)
		if err != nil{
			panic(err)
		}
		conn.Write(bytes)
	}
}
