#!/bin/sh

ENV PROTOC_ZIP=protoc-3.13.0-linux-x86_64.zip

apt-get update && apt-get install -y unzip

curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
    && unzip -o $PROTOC_ZIP -d /usr/local 'include/*' \
    && rm -f $PROTOC_ZI

go mod download
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
protoc --proto_path proto --go_out=:show  --go-grpc_out=:show --go_opt=paths=import proto/*.proto