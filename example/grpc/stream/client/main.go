package main

/**
* @Time       : 2020/11/22 9:08 下午
* @Author     : xumamba
* @Description: stream example client
 */

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "go-life/example/grpc/stream"
	gp "go-life/library/grpcpool"
)

func getStudent(client pb.StudentsClient, student *pb.Student) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	feature, err := client.GetStudent(ctx, student)
	if err != nil {
		log.Fatalf("%v.GetStudent(_) = _, %v", client, err)
	}
	log.Println(feature)
}

func listStudents(client pb.StudentsClient, class *pb.Class) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	studentsStream, err := client.ListStudents(ctx, class)
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	for {
		student, err := studentsStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Printf("Student[id: %d, name: %s]", student.GetId(), student.GetName())
	}
}

func recordStudents(client pb.StudentsClient) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	stream, err := client.RecordStudents(ctx)
	if err != nil {
		log.Fatalf("%v.RecordStudents(_) = _, %v", client, err)
	}

	for i := 1; i <= 80; i++ {
		err := stream.Send(&pb.Student{
			Id: int64(i),
			Name: fmt.Sprintf("student%d", i)})
		if err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, i, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("RecordStudents: %+v", reply)
}


func main() {
	conn, err := gp.GetConn("localhost:50005")
	if err != nil {
		log.Fatalf("failed get connection from pool: %v", err)
	}
	client := pb.NewStudentsClient(conn)

	recordStudents(client)

	getStudent(client, &pb.Student{Id: 1, Name: "student1"})

	listStudents(client, &pb.Class{ClassId: 1})
}
