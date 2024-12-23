package dto

type GameDTO struct {
	Name           string  `json:"name" binding:"required"`
	Description    string  `json:"description"`
	Size           float64 `json:"size"`
	Price          float64 `json:"price" binding:"required"`
	AgeRestriction string  `json:"age_restriction"`
	DeveloperId    int     `json:"developer_id" binding:"required"`
	CategoryIds    []int   `json:"category_ids"`
	RestrictIds    []int   `json:"restrict_ids"`
}

type GameUpdateDTO struct {
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Size           float64 `json:"size"`
	Price          float64 `json:"price"`
	AgeRestriction string  `json:"age_restriction"`
	DeveloperId    int     `json:"developer_id"`
	CategoryIds    []int   `json:"category_ids"`
	RestrictIds    []int   `json:"restrict_ids"`
}

type DeveloperDTO struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Description string `json:"description" binding:"required"`
	Country     string `json:"country" binding:"required"`
}

type DeveloperUpdateDTO struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Country     string `json:"country"`
}

type CategoryDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CategoryUpdateDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RestrictDTO struct {
	GameID int    `json:"game_id" binding:"required"`
	Region string `json:"region" binding:"required"`
}

type RestrictUpdateDTO struct {
	GameID int    `json:"game_id"`
	Region string `json:"region"`
}
