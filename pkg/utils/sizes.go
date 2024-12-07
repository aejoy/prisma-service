package utils

import (
	"image"
	_ "image/jpeg" //nolint:nolintlint
	_ "image/png"
	"io"

	"github.com/aejoy/prisma-service/pkg/consts"
	"github.com/aejoy/prisma-service/pkg/types"
	_ "github.com/gen2brain/avif" //nolint:nolintlint
	"github.com/pkg/errors"
)

func GetImageDimensions(typ types.ImageType, src io.Reader) (height, width int, err error) {
	switch typ {
	case consts.AvatarImageType:
		height = consts.AvatarHeight
		width = consts.AvatarWidth
	case consts.BannerImageType:
		height = consts.BannerHeight
		width = consts.BannerWidth
	case consts.PhotoImageType:
		img, _, imcodeErr := image.Decode(src)
		if imcodeErr != nil {
			return height, width, errors.Wrap(err, "Image Decode")
		}

		b := img.Bounds()

		height = b.Dy()
		width = b.Dx()
		err = imcodeErr
	}

	return height, width, err
}
