package photo

import (
	"bytes"
	"github.com/aejoy/prisma-service/internal/interfaces"
	"github.com/aejoy/prisma-service/internal/models"
	"github.com/aejoy/prisma-service/pkg/consts"
	"github.com/aejoy/prisma-service/pkg/converter"
	"github.com/aejoy/prisma-service/pkg/utils"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/pkg/errors"
	"mime/multipart"
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

func (s *Service) SavePhoto(creator string, src multipart.File, height, width int) (models.Photo, error) {
	photoID, err := gonanoid.Generate(consts.NanoAlphabet, consts.NanoLength)
	if err != nil {
		return models.Photo{}, errors.Wrap(err, "gonanoid.Generate")
	}

	if _, err := src.Seek(0, 0); err != nil {
		return models.Photo{}, errors.Wrap(err, "src.Seek(1)")
	}

	img, err := converter.ToAVIF(src, height, width)
	if err != nil {
		return models.Photo{}, errors.Wrap(err, "converter.ToAVIF")
	}

	size := len(img)

	if _, err := src.Seek(0, 0); err != nil {
		return models.Photo{}, errors.Wrap(err, "src.Seek(2)")
	}

	url, err := s.storage.SaveObject(utils.GetTimestampedPath(photoID), bytes.NewReader(img))
	if err != nil {
		return models.Photo{}, errors.Wrap(err, "storage.SaveObject")
	}

	if _, err := src.Seek(0, 0); err != nil {
		return models.Photo{}, errors.Wrap(err, "src.Seek(3)")
	}

	blurHash, err := utils.GetBlurHash(img)
	if err != nil {
		return models.Photo{}, errors.Wrap(err, "utils.GetBlurHash")
	}

	if err := s.repository.CreatePhoto(photoID, creator, url, blurHash, height, width, size); err != nil {
		return models.Photo{}, errors.Wrap(err, "repository.CreatePhoto")
	}

	return models.Photo{
		ID:       photoID,
		URL:      url,
		BlurHash: blurHash,
		Height:   height,
		Width:    width,
		Size:     size,
	}, nil
}
