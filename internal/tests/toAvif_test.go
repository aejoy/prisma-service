package tests

import (
	"bytes"
	"github.com/aejoy/prisma-service/pkg/photos"
	"io"
	"os"
	"testing"
)

func TestToAVIF(t *testing.T) {
	t.Run("avatar", func(t *testing.T) {
		in, err := os.Open("./orig/ap.png")
		if err != nil {
			t.Error(err)
		}

		avifImg, err := photos.ProcessAvatar(in)
		if err != nil {
			t.Error(err)
		}

		r, err := io.ReadAll(bytes.NewBuffer(avifImg))
		if err != nil {
			t.Error(err)
		}

		if err := os.WriteFile("./avif/ap.avif", r, 0600); err != nil {
			t.Error(err)
		}
	})

	t.Run("banner", func(t *testing.T) {
		in, err := os.Open("./orig/0aea769027635bfe5d6e8f0c6bf72102.png")
		if err != nil {
			t.Error(err)
		}

		avifImg, err := photos.ProcessBanner(in)
		if err != nil {
			t.Error(err)
		}

		r, err := io.ReadAll(bytes.NewBuffer(avifImg))
		if err != nil {
			t.Error(err)
		}

		if err := os.WriteFile("./avif/0aea769027635bfe5d6e8f0c6bf72102.avif", r, 0600); err != nil {
			t.Error(err)
		}
	})

	t.Run("photo", func(t *testing.T) {
		in, err := os.Open("./orig/ee5d4adf8b65babba897c044cae8fbd8.jpg")
		if err != nil {
			t.Error(err)
		}

		height, width, err := photos.GetMediaSizes(in)
		if err != nil {
			t.Error(err)
		}

		if _, err := in.Seek(0, 0); err != nil {
			t.Fatal(err)
		}

		img, err := photos.ToAVIF(in, height, width)
		if err != nil {
			t.Error(err)
		}

		if err := os.WriteFile("./avif/ee5d4adf8b65babba897c044cae8fbd8.avif", img, 0600); err != nil {
			t.Error(err)
		}
	})
}
