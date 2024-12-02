package http

import (
	"bytes"
	"fmt"
	"github.com/aejoy/prisma-service/internal/handlers/http/dto"
	"github.com/aejoy/prisma-service/internal/models"
	"github.com/aejoy/prisma-service/pkg/photos"
	"github.com/aejoy/prisma-service/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *Handlers) Upload(ctx fiber.Ctx) error {
	res := dto.Photos{}

	defer func() {
		utils.ReturnFiberResponse(ctx, res)
	}()

	if avatar, err := ctx.FormFile("avatar"); avatar != nil {
		if err != nil {
			if err.Error() != "there is no uploaded file associated with the given key" {
				return err
			}
		} else {
			file, err := avatar.Open()
			if err != nil {
				res.ErrorMessage = err.Error()
				return nil
			}

			img, err := photos.ProcessAvatar(file)
			if err != nil {
				res.ErrorMessage = err.Error()
				return nil
			}

			size := len(img)

			id, url, blurHash, err := h.prismaService.CreatePhoto("", bytes.NewReader(img), 256, 256, size)
			if err != nil {
				res.ErrorMessage = err.Error()
				return nil
			}

			res.Avatar = &models.Photo{ID: id, URL: url, BlurHash: blurHash, Height: 256, Width: 256, Size: size}
		}
	}

	if banner, err := ctx.FormFile("banner"); banner != nil {
		if err != nil {
			if err.Error() != "there is no uploaded file associated with the given key" {
				return err
			}
		} else {
			file, err := banner.Open()
			if err != nil {
				res.ErrorMessage = err.Error()
				return nil
			}

			img, err := photos.ProcessBanner(file)
			if err != nil {
				res.ErrorMessage = err.Error()
				return nil
			}

			if _, err := file.Seek(0, 0); err != nil {
				res.ErrorMessage = err.Error()
				return nil
			}

			size := len(img)

			id, url, blurHash, err := h.prismaService.CreatePhoto("", bytes.NewReader(img), 382, 728, size)
			if err != nil {
				res.ErrorMessage = err.Error()
				return nil
			}

			res.Banner = &models.Photo{ID: id, URL: url, BlurHash: blurHash, Height: 382, Width: 728, Size: size}
		}
	}

	for i := 0; i <= 5; i++ {
		photo, err := ctx.FormFile(fmt.Sprintf("photo%d", i+1))
		if err != nil {
			if err.Error() == "there is no uploaded file associated with the given key" {
				break
			} else {
				res.ErrorMessage = err.Error()
				return nil
			}
		}

		file, err := photo.Open()
		if err != nil {
			res.ErrorMessage = err.Error()
			return nil
		}

		height, width, err := photos.GetMediaSizes(file)
		if err != nil {
			res.ErrorMessage = err.Error()
			return nil
		}

		if _, err := file.Seek(0, 0); err != nil {
			res.ErrorMessage = err.Error()
			return nil
		}

		img, err := photos.ToAVIF(file, height, width)
		if err != nil {
			res.ErrorMessage = err.Error()
			return nil
		}

		size := len(img)

		id, url, blurHash, err := h.prismaService.CreatePhoto("", bytes.NewReader(img), height, width, size)
		if err != nil {
			res.ErrorMessage = err.Error()
			return nil
		}

		res.Photos = append(res.Photos, &models.Photo{ID: id, URL: url, BlurHash: blurHash, Height: height, Width: width, Size: size})
	}

	return nil
}
