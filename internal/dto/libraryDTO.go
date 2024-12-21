package dto

import (
	"TIPPr4/internal/models"
)

type LibraryDTO struct {
	ID     int           `json:"id"`
	UserID int           `json:"user_id"`
	Games  []models.Game `json:"games"`
}
