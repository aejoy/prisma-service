package dto

import "github.com/aejoy/prisma-service/internal/models"

type Photos struct {
	ErrorMessage string          `json:"error_message,omitempty"`
	Photo        *models.Photo   `json:"photo,omitempty"`
	Photos       []*models.Photo `json:"photos,omitempty"`
}
