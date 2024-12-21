package dto

import "TIPPr4/internal/models"

type CartDTO struct {
	ID     int           `json:"id"`
	UserID int           `json:"user_id"`
	Games  []models.Game `json:"games"`
}
