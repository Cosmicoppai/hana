package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"hana/show"
)

type VideoServer struct {
	show.UnimplementedVideoServiceServer
}

func NewVideoServer() *VideoServer {
	return &VideoServer{}
}

func (server *VideoServer) GetVideoMetadata(context.Context, *show.VideoId) (*show.VideoMetaData, error) {

}
func (server *VideoServer) GetVideosMetadata(context.Context, *emptypb.Empty) (*show.VideosMetaData, error) {

}
func (server *VideoServer) VideoMetadataByShow(context.Context, *show.ShowId) (*show.VideosMetaData, error) {

}

func (server *VideoServer) GetVideo(*show.VideoId, show.VideoService_GetVideoServer) error {

}

func (server *VideoServer) AddVideo(context.Context, *show.UploadVideo) (*show.Success, error) {

}
