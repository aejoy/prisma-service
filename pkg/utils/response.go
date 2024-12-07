package utils

import (
	"encoding/json"
	"net/http"

	"github.com/aejoy/prisma-service/internal/handlers/http/dto"
	"github.com/gofiber/fiber/v3"
)

func ReturnFiberResponse(ctx fiber.Ctx, res dto.Photos) {
	body, err := json.Marshal(res)
	if err != nil {
		res.ErrorMessage = err.Error()
	}

	if res.ErrorMessage != "" {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.Set("Content-Type", "application/json")

	_ = ctx.Send(body)
}
