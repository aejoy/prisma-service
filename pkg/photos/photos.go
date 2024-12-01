package photos

import (
	"github.com/davidbyttow/govips/v2/vips"
	_ "github.com/gen2brain/avif"
	"github.com/pkg/errors"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

func ProcessAvatar(file io.Reader) ([]byte, error) {
	img, err := ToAVIF(file, 256, 256)
	return img, err
}

func ProcessBanner(file io.Reader) ([]byte, error) {
	img, err := ToAVIF(file, 382, 728)
	return img, err
}

func ToAVIF(in io.Reader, height, width int) ([]byte, error) {
	imgRef, err := vips.NewImageFromReader(in)
	if err != nil {
		return nil, errors.Wrap(err, "LoadImageFromBuffer")
	}

	if err := imgRef.Thumbnail(width, height, vips.InterestingCentre); err != nil {
		return nil, err
	}

	defer imgRef.Close()

	img, _, err := imgRef.ExportAvif(&vips.AvifExportParams{
		Quality:  85,
		Lossless: false,
		Bitdepth: 12,
	})
	if err != nil {
		return nil, errors.Wrap(err, "ExportAVIF")
	}

	return img, nil
}

func GetMediaSizes(file io.Reader) (height, width int, err error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return height, width, errors.Wrap(err, "Image Decode")
	}

	b := img.Bounds()
	return b.Dy(), b.Dx(), err
}
