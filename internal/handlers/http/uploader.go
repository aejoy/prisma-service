package http

import (
	"mime/multipart"

	"github.com/aejoy/prisma-service/internal/handlers/http/dto"
	"github.com/aejoy/prisma-service/pkg/consts"
	"github.com/aejoy/prisma-service/pkg/types"
	"github.com/aejoy/prisma-service/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

type UploadTask struct {
	Type   types.ImageType
	Header *multipart.FileHeader
}

func (h *Handlers) Upload(ctx fiber.Ctx) error {
	res := dto.Photos{}

	defer func() {
		utils.ReturnFiberResponse(ctx, res)
	}()

	var typ types.ImageType

	switch ctx.FormValue("type", "photo") {
	case "avatar":
		typ = consts.AvatarImageType
	case "banner":
		typ = consts.BannerImageType
	case "photo":
		typ = consts.PhotoImageType
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		res.ErrorMessage = err.Error()
		return nil
	}

	src, err := file.Open()
	if err != nil {
		res.ErrorMessage = err.Error()
		return nil
	}

	height, width, err := utils.GetImageDimensions(typ, src)
	if err != nil {
		res.ErrorMessage = err.Error()
		return nil
	}

	photo, err := h.prismaService.SavePhoto("", src, height, width)
	if err != nil {
		res.ErrorMessage = err.Error()
		return nil
	}

	if err := src.Close(); err != nil {
		res.ErrorMessage = err.Error()
		return nil
	}

	res.Photo = &photo

	return nil
}
