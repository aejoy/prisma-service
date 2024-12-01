package photos

import (
	"bytes"
	"fmt"
	"github.com/aejoy/prisma-service/internal/interfaces"
	"github.com/aejoy/prisma-service/internal/models"
	"github.com/aejoy/prisma-service/pkg/consts"
	"github.com/buckket/go-blurhash"
	"github.com/gen2brain/avif"
	"github.com/matoous/go-nanoid/v2"
	"github.com/pkg/errors"
	"io"
	"time"
)

type Service struct {
	repository interfaces.Repository
	storage    interfaces.Storage
}

func NewPhotoService(repo interfaces.Repository, storage interfaces.Storage) *Service {
	return &Service{repo, storage}
}

func (s *Service) GetPhotos() ([]*models.Photo, error) {
	return s.repository.GetPhotos()
}

func (s *Service) GetPhotosByIDs(photoIDs []string) ([]*models.Photo, error) {
	return s.repository.GetPhotosByIDs(photoIDs)
}

func (s *Service) CreatePhoto(creator string, file io.Reader, height, width, size int) (string, string, string, error) {
	photoID, err := gonanoid.Generate(consts.NanoAlphabet, 16)
	if err != nil {
		return photoID, "", "", errors.Wrap(err, "nanoid.generate")
	}

	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	mediaURL := fmt.Sprintf("%d/%02d/%02d/%s.avif", year, int(month), day, photoID)

	filePtr, err := io.ReadAll(file)
	if err != nil {
		return photoID, "", "", errors.Wrap(err, "io.readAll")
	}

	url, err := s.storage.SaveObject(mediaURL, bytes.NewReader(filePtr))
	if err != nil {
		return photoID, url, "", errors.Wrap(err, "storage.saveObject")
	}

	img, err := avif.Decode(bytes.NewReader(filePtr))
	if err != nil {
		return photoID, url, "", errors.Wrap(err, "avif.decode")
	}

	blurHash, err := blurhash.Encode(4, 3, img)
	if err != nil {
		return photoID, url, blurHash, errors.Wrap(err, "blurhash.encode")
	}

	return photoID, url, blurHash, s.repository.CreatePhoto(photoID, creator, url, blurHash, height, width, size)
}
