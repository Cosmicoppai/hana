#!/bin/sh

apt-get update && apt-get install -y unzip

apt install -y protobuf-compiler

go mod tidy

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --proto_path proto --go_out=:show  --go-grpc_out=:show --go_opt=paths=import proto/*.proto