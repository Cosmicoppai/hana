package service

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hana/show"
	"log"
)

const maxVideoSize = 1 << 29

type VideoServer struct {
	show.UnimplementedVideoServiceServer
	store VideoStore
}

func NewVideoServer(store VideoStore) *VideoServer {
	return &VideoServer{store: store}
}

func (server *VideoServer) GetVideoMetadata(context.Context, *show.VideoId) (*show.VideoMetaData, error) {
	return nil, nil

}
func (server *VideoServer) GetVideosMetadata(context.Context, *emptypb.Empty) (*show.VideosMetaData, error) {
	return nil, nil

}

func (server *VideoServer) GetVideo(*show.VideoId, show.VideoService_GetVideoServer) error {
	return nil
}

func (server *VideoServer) AddVideo(stream show.VideoService_AddVideoServer) error {
	req, err := stream.Recv()
	if err != nil {
		log.Println(err)
		return status.Errorf(codes.Unknown, "cannot receive the video info")
	}

	poster := bytes.Buffer{}
	sub := bytes.Buffer{}
	_poster := req.GetMetaData().Poster
	_sub := req.GetMetaData().Sub

	poster.Write(_poster)
	if _sub != nil {
		sub.Write(_sub)
	}
	log.Println("Received poster and sub")

	//videoData := bytes.Buffer{}
	//videoSize := 0
	//for {
	//	log.Println("Receiving video data ")
	//	req, err := stream.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		return status.Errorf(codes.Unknown, "cannot receive video")
	//	}
	//	chunk := req.GetVideo().GetVideo()
	//	size := len(chunk)
	//	videoSize += size
	//	if videoSize > maxVideoSize {
	//		return status.Errorf(codes.InvalidArgument, "video is too large %d > %d", videoSize, maxVideoSize)
	//	}
	//	videoData.Write(chunk)
	//}
	//videoId, err := server.store.Save("mp4", videoData)
	//if err != nil {
	//	log.Println(err)
	//	return status.Errorf(codes.Internal, "Cannot save video")
	//}
	videoId := uuid.New()
	err = server.store.SaveMetaData(videoId, poster, sub)
	if err != nil {
		log.Println(err)
		return status.Errorf(codes.Internal, "Cannot save video meta data")
	}
	res := &show.VideoId{
		Id: videoId.String(),
	}
	err = stream.SendAndClose(res)
	if err != nil {
		return status.Errorf(codes.Internal, "Cannot send response")
	}
	log.Println("Data saved successfully")
	return nil
}
