package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func GetPaginationParams(ctx fiber.Ctx) (offset int, count int, err error) {
	if offsetParam := ctx.Query("offset"); offsetParam != "" {
		offset, err = strconv.Atoi(offsetParam)
		if err != nil {
			return offset, count, err
		}
	}

	if countParam := ctx.Query("count"); countParam != "" {
		count, err = strconv.Atoi(countParam)
		if err != nil {
			return offset, count, err
		}
	}

	if count <= 0 {
		count = 50
	}

	return offset, count, nil
}
