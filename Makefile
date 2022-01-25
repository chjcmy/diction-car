# Makefile

init:
	go get github.com/golang/protobuf/protoc-gen-go
	go get google.golang.org/grpc
	go mod download

car: ./proto/car.proto
	protoc ./proto/car.proto --go_out=./pb --go-grpc_out=./pb

diction: ./proto/diction.proto
	protoc ./proto/diction.proto --go_out=./pb --go-grpc_out=./pb