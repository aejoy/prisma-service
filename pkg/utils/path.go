package utils

import (
	"fmt"
	"time"
)

func GetTimestampedPath(photoID string) string {
	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()

	return fmt.Sprintf("%d/%02d/%02d/%s.avif", year, int(month), day, photoID)
}
