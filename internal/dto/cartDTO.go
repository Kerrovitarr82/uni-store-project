package dto

import "uniStore/internal/models"

type CartDTO struct {
	ID     int           `json:"id"`
	UserID int           `json:"user_id"`
	Games  []models.Game `json:"games"`
}
