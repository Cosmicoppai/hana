package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hana/show"
	"log"
	"os"
	"time"
)

func testAddVideo(videoClient show.VideoServiceClient) {
	addVideo(videoClient, "./video/test.mp4", "./poster/m.png")
}

func addVideo(client show.VideoServiceClient, videoPath string, posterPath string) {
	log.Println("Opening poster file...")
	posterFile, err := os.Open(posterPath)
	if err != nil {
		log.Println("Error in opening the file", err)
	}
	info, _ := posterFile.Stat()
	posterSize := info.Size()
	poster := make([]byte, posterSize)
	log.Println("Reading poster...")
	_, err = posterFile.Read(poster)
	if err != nil {
		log.Println(err)
	}
	metaReq := &show.UploadVideo{Data: &show.UploadVideo_MetaData{MetaData: &show.MetaData{Poster: poster}}}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	stream, err := client.AddVideo(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	err = stream.Send(metaReq)
	if err != nil {
		log.Fatalln("Unable to send metaData", err)
	}
	id, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("Error while closing stream", err)
	}
	log.Println("Received id: ", id)

}

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
	videoClient := show.NewVideoServiceClient(conn)
	testAddVideo(videoClient)
}
