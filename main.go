package main

import (
	"fmt"
	"grpc/test/server"
)

func main() {
	fmt.Println("start main")
	server.Handler()
}
