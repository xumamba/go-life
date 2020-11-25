package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	pb "go-life/example/grpc/stream"
)

type StudentServer struct {
	pb.UnimplementedStudentsServer

	locker    sync.RWMutex
	classInfo map[int32][]*pb.Student
}

func (s *StudentServer) GetStudent(ctx context.Context, student *pb.Student) (*pb.Feature, error) {
	log.Printf("[GetStudent] receive request: %+v\n", student)

	classID := s.divideClass(student.Id)
	s.locker.RLock()
	defer s.locker.RUnlock()
	for _, body := range s.classInfo[classID] {
		if body.Id == student.Id {
			return &pb.Feature{ClassId: classID, StudentInfo: student}, nil
		}
	}
	return nil, errors.New("no information about the student was found")
}

func (s *StudentServer) ListStudents(class *pb.Class, studentsServer pb.Students_ListStudentsServer) error {
	log.Printf("[ListStudents] receive request: %+v\n", class)
	s.locker.RLock()
	defer s.locker.RUnlock()
	for _, student := range s.classInfo[class.ClassId] {
		log.Println(student)
		if err := studentsServer.Send(student); err != nil {
			return err
		}
	}
	return nil
}

func (s *StudentServer) RecordStudents(studentsServer pb.Students_RecordStudentsServer) error {
	var studentCount, classCount int32
	startTime := time.Now()
	s.locker.Lock()
	defer s.locker.Unlock()
	for {
		student, err := studentsServer.Recv()
		log.Printf("[RecordStudents] receive student: %+v\n", student)
		if err == io.EOF {
			endTime := time.Now()
			return studentsServer.SendAndClose(&pb.StudentSummary{
				StudentCount: studentCount,
				ClassCount:   classCount,
				ElapsedTime:  int32(endTime.Sub(startTime).Microseconds()),
			})
		}
		if err != nil {
			return err
		}
		studentCount++
		classID := s.divideClass(student.Id)
		if _, ok := s.classInfo[classID]; !ok {
			classCount++
			s.classInfo[classID] = make([]*pb.Student, 0)
		}
		s.classInfo[classID] = append(s.classInfo[classID], student)
	}
}

func NewServer() *StudentServer {
	s := &StudentServer{
		classInfo: make(map[int32][]*pb.Student),
	}
	return s
}

func (s *StudentServer) divideClass(id int64) int32 {
	return int32(id % 8)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:50005")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStudentsServer(s, NewServer())
	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
