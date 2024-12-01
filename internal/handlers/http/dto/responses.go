package dto

import "github.com/aejoy/prisma-service/internal/models"

type Photos struct {
	ErrorMessage string          `json:"error_message,omitempty"`
	Photos       []*models.Photo `json:"photos,omitempty"`
	Avatar       *models.Photo   `json:"avatar,omitempty"`
	Banner       *models.Photo   `json:"banner,omitempty"`
}
