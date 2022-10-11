package server

import (
	"context"
	pb "grpc/test/grpc"
	"grpc/test/rdb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	// luôn phải có trường này
	pb.UnimplementedTestServiceServer
	Conn rdb.Connection
}

// GRPC bắt phải có hàm đã được định nghĩa ở trong file pb.go
func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{Result: in.GetNum1() + in.GetNum2()}, nil
}

func (s *server) ListPostInvitee(ctx context.Context, in *pb.ListPostInviteeRequest) (*pb.ListPostInviteeResponse, error) {
	listInvitee, err := s.Conn.ListPostInvitee(ctx, in.PostId)
	if err != nil {
		return nil, err
	}
	return &pb.ListPostInviteeResponse{
		InviteeList: ListPostInviteeToGRPC(listInvitee),
	}, nil
}

// Convert function
func ListPostInviteeToGRPC(lpi []rdb.PostInvitee) []*pb.PostInviteeResponse {
	vv := make([]*pb.PostInviteeResponse, len(lpi))
	for i, pi := range lpi {
		vv[i] = PostInviteeToGRPC(pi)
	}
	return vv
}

func PostInviteeToGRPC(pi rdb.PostInvitee) *pb.PostInviteeResponse {
	return &pb.PostInviteeResponse{
		Id:               uint32(pi.User.ID),
		EmUserId:         uint32(pi.User.EmUserID),
		Name:             pi.User.Name,
		Email:            pi.User.Email,
		IconId:           uint32(pi.User.IconID),
		UserType:         pb.UserType(pi.User.UserType),
		PostInviteeToken: pi.UUID.UUID.String(),
		Comment:          *pi.Comment,
		Passcode:         *pi.Passcode,
	}
}

func Handler() {
	port := "9091"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// connection to RDB
	conn, err := rdb.Connect()
	if err != nil {
		log.Fatalf("Cannot connect to RDB: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTestServiceServer(s, &server{Conn: conn})

	log.Printf("Listening on %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
