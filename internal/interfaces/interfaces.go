package interfaces

import (
	"io"
	"mime/multipart"

	"github.com/aejoy/prisma-service/internal/models"
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
	SavePhoto(creator string, src multipart.File, height, width int) (models.Photo, error)
}
