package server

import (
	"context"
	"fmt"
	pb "grpc/test/grpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTestServiceServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Println("AAAA")
	return &pb.SumResponse{Result: in.GetNum1() + in.GetNum2()}, nil
}

func Handler() {
	port := "9091"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err) // TODO 問題点 No.46
	}

	s := grpc.NewServer()
	pb.RegisterTestServiceServer(s, &server{})

	log.Printf("Listening on %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
