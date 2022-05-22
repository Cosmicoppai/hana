## Introduction
<img align="right" width="120" height="120" src="logo.png" alt="icon">


Video Streaming over internet using gRPC.
The client uploads the video and video-metadata to the server using Client-Streaming.
The server process and saves the video and its metadata for future streaming.



## Requirements

* If you don't want to install dependencies on your machine skip to [Installation using Docker](#installation-using-docker).

Install [Golang](https://go.dev/doc/install)

Install Protoc ```apt install -y protobuf-compiler```

Install Makefile ```apt install make```


## Installation

```bash
go mod download
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
Make gen // convert .proto files into go files
Make server // will start the streaming server on 0.0.0.0:8096
Make Client // will start the client
```

## Installation using Docker

```bash
docker-compose up -d
```
Will start the Streaming Server on [localhost:8096](http://localhost:8096)

To start Client

```bash
docker exec -it sh hana go cmd/client/main.go --address 0.0.0.0:8096
```

## Contributing
This Project is open source, feel free to raise any issue/question or make PR :)

## License

Source Code is licensed under [MIT LICENSE](./LICENSE)


## TO DO
- #### Refactoring
- #### Add more tests
- #### Client Authentication
