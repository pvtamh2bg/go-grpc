gen-user-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/test.proto

run-server:
	go run main.go

run-client:
	go run client/client.go
