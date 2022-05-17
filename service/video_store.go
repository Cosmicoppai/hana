package service

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type VideoStore interface {
	Save(videoName string, videoTyp string, videoData bytes.Buffer) (uuid.UUID, error)
	SaveMetaData(videoName string, videoId uuid.UUID, poster bytes.Buffer, sub bytes.Buffer) error
}

type DiskVideoStore struct {
	mutex       sync.RWMutex
	videoFolder string
	videos      map[string]*VideoInfo
}

type VideoInfo struct {
	Type string
	Path string
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
		Path: filePath,
		Type: videoTyp,
	}
	return id, nil
}

func (store *DiskVideoStore) SaveMetaData(videoName string, videoId uuid.UUID, poster bytes.Buffer, sub bytes.Buffer) error {
	if poster.Len() == 0 {
		return status.Errorf(codes.InvalidArgument, "send poster in proper format")
	}
	dirPath := filepath.Join(store.videoFolder, videoId.String())
	filePath := fmt.Sprintf(filepath.Join(dirPath, fmt.Sprintf("%s%s", videoName, ".png")))
	file, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		return status.Errorf(codes.Internal, "Internal server error")
	}
	_, err = poster.WriteTo(file)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal server error")
	}
	log.Println("Meta Data saved Successfully")
	return nil
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
