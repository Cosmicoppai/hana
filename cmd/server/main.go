package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"hana/service"
	"hana/show"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 9000, "Server Port")
	flag.Parse()
	log.Printf("dialling on port %d", *port)

	address := fmt.Sprintf("0.0.0.0:%d", *port)

	gRPCServer := grpc.NewServer()

	videoStore := service.NewDiskStore("./video") // create an InMem VideoStore

	videoServer := service.NewVideoServer(videoStore)
	show.RegisterVideoServiceServer(gRPCServer, videoServer)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Cannot start server: ", err)
	}

	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
