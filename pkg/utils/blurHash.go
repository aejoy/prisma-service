package utils

import (
	"bytes"
	"image"
	_ "image/jpeg" //nolint:nolintlint
	_ "image/png"

	"github.com/aejoy/prisma-service/pkg/consts"
	"github.com/buckket/go-blurhash"
	_ "github.com/gen2brain/avif" //nolint:nolintlint
	"github.com/pkg/errors"
)

func GetBlurHash(src []byte) (blurHash string, err error) {
	img, _, err := image.Decode(bytes.NewReader(src))
	if err != nil {
		return blurHash, errors.Wrap(err, "avif.decode")
	}

	blurHash, err = blurhash.Encode(consts.BlurHashXComponents, consts.BlurHashYComponents, img)

	return blurHash, errors.Wrap(err, "blurHash.encode")
}
