package dto

import (
	"time"
	"uniStore/internal/models"
)

type OrderDTO struct {
	ID        int           `json:"id"`
	UserID    int           `json:"user_id"`
	Games     []models.Game `json:"games"`
	TotalCost float64       `json:"total_cost"`
	CreatedAt time.Time     `json:"created_at"`
}
