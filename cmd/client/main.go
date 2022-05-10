package main

import (
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	servAddress := flag.String("address", "", "Server Address")
	flag.Parse()
	log.Printf("dialling server at %s", *servAddress)

	conn, err := grpc.Dial(*servAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)
	if err != nil {
		log.Fatalln(err)
	}
	/*
		videoClient := service.NewVideoClient
		create a request object
		make request using videoClient
	*/
}
