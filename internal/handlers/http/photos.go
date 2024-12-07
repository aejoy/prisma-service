package http

import (
	"strings"

	"github.com/aejoy/prisma-service/internal/handlers/http/dto"
	"github.com/aejoy/prisma-service/internal/models"
	"github.com/aejoy/prisma-service/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *Handlers) Photos(ctx fiber.Ctx) error {
	res := dto.Photos{}

	defer func() {
		utils.ReturnFiberResponse(ctx, res)
	}()

	var photos []*models.Photo

	if ids := ctx.Query("ids"); ids != "" {
		photosByIDs, err := h.prismaService.GetPhotosByIDs(strings.Split(ids, ","))
		if err != nil {
			res.ErrorMessage = err.Error()
			return nil
		}

		photos = photosByIDs
	} else {
		offset, count, err := utils.GetPaginationParams(ctx)
		if err != nil {
			res.ErrorMessage = err.Error()
			return nil
		}

		allPhotos, err := h.prismaService.GetPhotos(offset, count)
		if err != nil {
			res.ErrorMessage = err.Error()
			return nil
		}

		photos = allPhotos
	}

	res.Photos = make([]*models.Photo, 0, len(photos))
	res.Photos = photos

	return nil
}
