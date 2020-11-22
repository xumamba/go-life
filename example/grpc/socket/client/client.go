package main

/**
 * @DateTime   : 2020/11/18
 * @Author     : xumamba
 * @Description:
 **/

import (
	"bufio"
	"log"
	"net"
	"os"

	"go-life/example/grpc/socket"

	"github.com/gogo/protobuf/proto"
)

const serverAddr = "localhost:50051"

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("dail %s failed", serverAddr)
	}
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)

	buf := make([]byte, 4096, 4096)
	for scanner.Scan() {
		text := scanner.Text()
		request := &socket.HelloRequest{
			Greeting: text,
		}
		bytes, err := proto.Marshal(request)
		if err != nil {
			log.Fatalf("proto marshal failed:%v", err)
		}

		conn.Write(bytes)

		count, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		response := &socket.HelloResponse{}
		respData := buf[:count]
		err = proto.Unmarshal(respData, response)
		if err != nil {
			panic(err)
		}
		log.Printf("receive reply: %v", response.Reply)
		if text == "EOF" {
			return
		}
	}
}
