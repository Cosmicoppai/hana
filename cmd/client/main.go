package main

import (
	"bufio"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hana/show"
	"io"
	"log"
	"os"
	"time"
)

func testAddVideo(videoClient show.VideoServiceClient) {
	addVideo(videoClient, "./video/test.mp4", "./poster/m.png")
}

func addVideo(client show.VideoServiceClient, videoPath string, posterPath string) {
	maxMsgSize := 5243000
	maxSize := grpc.MaxCallSendMsgSize(maxMsgSize)
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
	metaReq := &show.UploadVideo{Data: &show.UploadVideo_MetaData{MetaData: &show.MetaData{Poster: poster,
		Name: "test"}}}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	stream, err := client.AddVideo(ctx, maxSize)
	if err != nil {
		log.Fatalln(err)
	}
	err = stream.Send(metaReq)
	if err != nil {
		log.Fatalln("Unable to send metaData", err)
	}

	file, err := os.Open(videoPath)
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 5242880)

	log.Println("Reading video data")
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Cannot read chunk to buffer: ", err)
		}
		req := &show.UploadVideo{Data: &show.UploadVideo_Video{Video: &show.Video{Video: buffer[:n]}}}

		err = stream.Send(req)
		if err != nil {
			log.Fatalln("Cannot send chunk to the server: ", err)
		}
	}

	id, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("Error while closing stream", err)
	}
	log.Println("Received id: ", id)
	receiveVideo(client, id)

}

func receiveVideo(client show.VideoServiceClient, videoId *show.VideoId) {
	maxMsgSize := 5243000
	maxSize := grpc.MaxCallRecvMsgSize(maxMsgSize)
	id := &show.VideoId{Id: videoId.Id}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	clientStream, err := client.GetVideo(ctx, id, maxSize)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("./client_req/test.mp4")
	if err != nil {
		log.Fatalln("error while creating file to receive video", err)
	}
	for {
		req, err := clientStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		_, err = file.Write(req.Video)
		if err != nil {
			log.Fatalln("error while writing to file", err)
		}

	}
	log.Println("Video Successfully received")
	_ = clientStream.CloseSend()
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
