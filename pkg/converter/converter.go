package converter

import (
	"bytes"
	"github.com/disintegration/imaging"
	"image"
	_ "image/jpeg" //nolint:nolintlint
	_ "image/png"
	"io"

	"github.com/aejoy/prisma-service/pkg/consts"
	"github.com/gen2brain/avif"
	"github.com/pkg/errors"
)

func ToAVIF(src io.Reader, height, width int) ([]byte, error) {
	img, _, err := image.Decode(src)
	if err != nil {
		return nil, errors.Wrap(err, "image.Decode")
	}

	dst := imaging.Fill(img, width, height, imaging.Center, imaging.Lanczos)

	var buf bytes.Buffer

	if err := avif.Encode(&buf, dst, avif.Options{
		Quality: consts.AVIFQuality,
	}); err != nil {
		return nil, errors.Wrap(err, "avif.Encode")
	}

	return buf.Bytes(), nil
}
