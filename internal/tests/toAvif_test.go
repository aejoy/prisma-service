package tests_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/aejoy/prisma-service/pkg/utils"

	"github.com/gen2brain/avif"
	"github.com/stretchr/testify/assert"

	"github.com/aejoy/prisma-service/pkg/consts"

	"github.com/aejoy/prisma-service/pkg/converter"
)

func TestToAVIF(t *testing.T) {
	tests := []struct {
		name          string
		src, target   string
		height, width int
	}{
		{name: "avatar", src: "./src/activity-pub.png", height: consts.AvatarHeight, width: consts.AvatarWidth},
		{name: "banner", src: "./src/0aea769027635bfe5d6e8f0c6bf72102.png", height: consts.BannerHeight, width: consts.BannerWidth},
		{name: "default", src: "./src/ee5d4adf8b65babba897c044cae8fbd8.jpg", height: 0, width: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			src, err := os.Open(test.src)
			if err != nil {
				t.Error(err)
			}

			if test.height == 0 && test.width == 0 {
				height, width, err := utils.GetImageDimensions(consts.AvatarImageType, src)
				if err != nil {
					t.Error(err)
				}

				if _, err := src.Seek(0, 0); err != nil {
					t.Error(err)
				}

				test.height = height
				test.width = width
			}

			target, err := converter.ToAVIF(src, test.height, test.width)
			if err != nil {
				t.Error(err)
			}

			if _, err := src.Seek(0, 0); err != nil {
				t.Error(err)
			}

			srcBytes, err := io.ReadAll(src)
			if err != nil {
				t.Error(err)
			}

			_, err = avif.Decode(bytes.NewBuffer(target))
			if err != nil {
				t.Error(err)
			}

			assert.Greater(t, len(srcBytes), 0, "src len <= 0")
			assert.Greater(t, len(target), 0, "target len <= 0")
		})
	}
}
