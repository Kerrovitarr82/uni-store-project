package dto

type ReviewDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Rating      int    `json:"rating" binding:"required"`
}

type ReviewUpdateDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
