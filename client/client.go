package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc/test/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connect() pb.TestServiceClient {
	conn, err := grpc.Dial("localhost:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close() // comment cái này lại vì kết thúc hàm này nó sẽ close connection

	return pb.NewTestServiceClient(conn)
}

func callSum(c pb.TestServiceClient) {
	log.Println("Calling Sum")
	res, err := c.Sum(context.Background(), &pb.SumRequest{Num1: 3, Num2: 5})
	if err != nil {
		log.Println("SumError")
	}
	log.Printf("Sum: %v", res)
}

func GetInviteeList(c pb.TestServiceClient) {
	res, err := c.ListPostInvitee(context.Background(), &pb.ListPostInviteeRequest{CompanyId: 1, PostId: 1})
	if err != nil {
		log.Println("get invitee error")
	}
	fmt.Println(res)
	log.Printf("Postinvite: %v", res)
}

func main() {
	c := connect()
	callSum(c)
	GetInviteeList(c)
}
