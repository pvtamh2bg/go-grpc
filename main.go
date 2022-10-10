package main

import (
	"fmt"
	"grpc/test/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("start main")
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading env file err=%+v", err) // TODO 問題点 No.46
	}
	server.Handler()
}
