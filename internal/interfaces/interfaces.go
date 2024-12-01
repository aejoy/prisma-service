package interfaces

import (
	"github.com/aejoy/prisma-service/internal/models"
	"io"
)

type Storage interface {
	SaveObject(key string, body io.Reader) (string, error)
}

type Repository interface {
	GetPhotos() ([]*models.Photo, error)
	GetPhotosByIDs(photoIDs []string) ([]*models.Photo, error)
	CreatePhoto(photoID, creator, url, blurHash string, height, width, size int) error
}

type Service interface {
	GetPhotos() ([]*models.Photo, error)
	GetPhotosByIDs(photoIDs []string) ([]*models.Photo, error)
	CreatePhoto(creator string, file io.Reader, height, width, size int) (string, string, string, error)
}
