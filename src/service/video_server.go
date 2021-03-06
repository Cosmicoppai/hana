package service

import (
	"bufio"
	"bytes"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hana/show"
	"io"
	"log"
	"os"
)

const maxVideoSize = 1 << 29
const defaultOffset = 25

type VideoServer struct {
	show.UnimplementedVideoServiceServer
	store VideoStore
}

func NewVideoServer(store VideoStore) *VideoServer {
	return &VideoServer{store: store}
}

func (server *VideoServer) GetVideoMetadata(ctx context.Context, videoId *show.VideoId) (*show.MetaData, error) {
	id := videoId.GetId()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "video ID is required")
	}
	metaData, err := server.store.GetMetaData(id)
	if err != nil {
		log.Println("error while retrieving in meta-data", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return metaData, nil

}
func (server *VideoServer) GetVideosMetadata(ctx context.Context, request *show.VideosMetaDataRequest) (*show.VideosMetaData, error) {
	limit := request.Limit
	offset := request.Offset
	totalVideos := uint32(server.store.TotalVideos())
	if offset > totalVideos {
		offset = totalVideos
	}
	if limit < offset {
		return nil, status.Errorf(codes.InvalidArgument, "limit is less than offset")
	}
	if limit == 0 {
		limit = offset + defaultOffset
	}
	if limit > totalVideos {
		limit = totalVideos
	}
	videosMetaData, err := server.store.GetMetaDatas(limit, offset)
	if err != nil {
		log.Println("error while returning videosMetaData", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return videosMetaData, nil

}

func (server *VideoServer) GetVideo(videoId *show.VideoId, stream show.VideoService_GetVideoServer) error {
	info, err := server.store.Find(videoId.Id)
	if err != nil {
		log.Println(err)
		return status.Errorf(codes.NotFound, "video with id %s doesn't exist", videoId.Id)
	}
	file, err := os.Open(info.VideoPath)
	if err != nil {
		log.Println("Unable to open file", err)
		return status.Errorf(codes.Internal, "Internal server error")
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 5242880)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error while reading video file", err)
			return status.Errorf(codes.Internal, "Internal server error")
		}
		video := &show.Video{Video: buffer[:n]}
		err = stream.Send(video)
		if err != nil {
			log.Println("error while sending video", err)
			return status.Errorf(codes.Internal, "Internal server error")
		}
	}
	return status.Errorf(codes.OK, "")
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
	videoName := req.GetMetaData().Name

	poster.Write(_poster)
	if _sub != nil {
		sub.Write(_sub)
	}
	log.Println("Received poster and sub")

	videoData := bytes.Buffer{}
	videoSize := 0
	for {
		log.Println("Receiving video data ")
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive video")
		}
		chunk := req.GetVideo().GetVideo()
		size := len(chunk)
		videoSize += size
		if videoSize > maxVideoSize {
			return status.Errorf(codes.InvalidArgument, "video is too large %d > %d", videoSize, maxVideoSize)
		}
		videoData.Write(chunk)
	}
	videoId, err := server.store.Save(videoName, "mp4", videoData)
	if err != nil {
		log.Println(err)
		return status.Errorf(codes.Internal, "Cannot save video")
	}
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
