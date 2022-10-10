package main

import (
	"context"
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
	log.Println("Calling Sum") // dùng Println thay cho log.Fatal vì log.Fatal sẽ dừng chương trình luôn
	res, err := c.Sum(context.Background(), &pb.SumRequest{Num1: 3, Num2: 5})
	if err != nil {
		log.Fatalln("SumError") // dùng Println thay cho log.Fatal vì log.Fatal sẽ dừng chương trình luôn
	}
	log.Printf("Sum: %v", res)
}

func main() {
	c := connect()
	callSum(c)
}
