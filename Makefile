gen-user-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/test.proto

run-server:
	go run main.go

run-client:
	go run client/client.go

migrate:
	migrate -path DDL -database "postgresql://emz-rdb-test-username:emz-rdb-test-password@127.0.0.1:5434/emz-rdb-test?sslmode=disable" -verbose up

test:
	go test -v ./...