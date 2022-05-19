package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hana/show"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type VideoStore interface {
	Save(videoName string, videoTyp string, videoData bytes.Buffer) (uuid.UUID, error)
	SaveMetaData(videoName string, videoId uuid.UUID, poster bytes.Buffer, sub bytes.Buffer) error
	Find(videoId string) (*VideoInfo, error)
	GetMetaDatas() (*show.VideosMetaData, error)
	GetMetaData(videoId string) (*show.MetaData, error)
}

type DiskVideoStore struct {
	mutex       sync.RWMutex
	videoFolder string
	videos      map[string]*VideoInfo
}

type VideoInfo struct {
	videoName  string
	videoType  string
	PosterPath string
	SubPath    string
	CreatedAt  timestamppb.Timestamp
	VideoPath  string
}

func NewDiskStore(path string) *DiskVideoStore {
	return &DiskVideoStore{videoFolder: path, videos: make(map[string]*VideoInfo)}
}

func (store *DiskVideoStore) Save(videoName string, videoTyp string, videoData bytes.Buffer) (uuid.UUID, error) {
	id := uuid.New()

	dirPath := filepath.Join(store.videoFolder, id.String())
	dir, err := os.Stat(dirPath)
	if err != nil || !dir.IsDir() {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			log.Println(err)
			return uuid.Nil, status.Errorf(codes.Internal, "Internal server Error")
		}
	}

	filePath := fmt.Sprintf(filepath.Join(dirPath, fmt.Sprintf("%s.%s", videoName, videoTyp)))

	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Unable to create file...")
		return uuid.Nil, err
	}
	_, err = videoData.WriteTo(file)
	if err != nil {
		log.Println("Error while writing to file", err)
		return uuid.Nil, err
	}
	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.videos[id.String()] = &VideoInfo{
		videoName: videoName,
		VideoPath: filePath,
		videoType: videoTyp,
		CreatedAt: timestamppb.Timestamp{Seconds: int64(time.Now().Second())},
	}
	return id, nil
}

func (store *DiskVideoStore) SaveMetaData(videoName string, videoId uuid.UUID, poster bytes.Buffer, sub bytes.Buffer) error {
	if poster.Len() == 0 {
		return status.Errorf(codes.InvalidArgument, "send poster in proper format")
	}
	dirPath := filepath.Join(store.videoFolder, videoId.String())
	posterPath := fmt.Sprintf(filepath.Join(dirPath, fmt.Sprintf("%s%s", videoName, ".png")))
	posterFile, err := os.Create(posterPath)
	if err != nil {
		log.Println(err)
		return status.Errorf(codes.Internal, "Internal server error")
	}
	_, err = poster.WriteTo(posterFile)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal server error")
	}
	id := videoId.String()
	store.videos[id].PosterPath = posterPath

	if sub.Len() > 0 {
		subPath := fmt.Sprintf(filepath.Join(dirPath, fmt.Sprintf("%s%s", videoName, ".vtt")))
		subFile, err := os.Create(subPath)
		if err != nil {
			log.Println("error while creating subtitle file: ", err)
			return status.Errorf(codes.Internal, "Internal server error")
		}
		_, err = sub.WriteTo(subFile)
		if err != nil {
			log.Println("Error while writing bytes of subFile: ", err)
			return status.Errorf(codes.Internal, "Internal server error")
		}
		store.videos[id].SubPath = subPath
	}

	log.Println("Meta Data saved Successfully")
	return nil
}

func (store *DiskVideoStore) Find(videoId string) (*VideoInfo, error) {
	if info, ok := store.videos[videoId]; ok {
		return info, nil
	}
	return nil, errors.New("video doesn't exist")
}

func (store *DiskVideoStore) GetMetaDatas() (*show.VideosMetaData, error) {
	return &show.VideosMetaData{}, nil
}

func (store *DiskVideoStore) GetMetaData(videoId string) (*show.MetaData, error) {
	if videoInfo, ok := store.videos[videoId]; ok {
		poster, posterErr := os.ReadFile(videoInfo.PosterPath)
		var (
			sub    []byte
			subErr error
		)
		if videoInfo.SubPath != "" {
			sub, subErr = os.ReadFile(videoInfo.SubPath)
		}
		if posterErr != nil || subErr != nil {
			log.Println("error while reading poster data", posterErr, subErr)
			return &show.MetaData{}, status.Errorf(codes.Internal, "Internal Server Error")
		}
		metaData := show.MetaData{Poster: poster, Sub: sub, CreatedAt: &videoInfo.CreatedAt, Name: videoInfo.videoName}
		return &metaData, nil
	}
	return &show.MetaData{}, nil
}

type DBVideoStore struct {
	mutex       sync.RWMutex
	videoFolder string
	connString  string
	videos      map[string]*VideoInfo
}

func NewDBStore(connString string) *DBVideoStore {
	return &DBVideoStore{connString: connString}
}

func (store *DBVideoStore) Save(videoName string, videoTyp string, videoData bytes.Buffer) (uuid.UUID, error) {
	return uuid.Nil, nil

}

func (store *DBVideoStore) SaveMetaData(videoName string, videoId uuid.UUID, poster bytes.Buffer, sub bytes.Buffer) error {
	return nil
}

func (store *DBVideoStore) Find(videoId string) (*VideoInfo, error) {
	if info, ok := store.videos[videoId]; ok {
		return info, nil
	}
	return nil, errors.New("video doesn't exist")
}

func (store *DBVideoStore) GetMetaDatas() (*show.VideosMetaData, error) {
	return &show.VideosMetaData{}, nil
}

func (store *DBVideoStore) GetMetaData(videoId string) (*show.MetaData, error) {
	return &show.MetaData{}, nil
}
