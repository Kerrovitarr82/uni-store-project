package dto

type LoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupDTO struct {
	Name       string `json:"name" binding:"required"`
	SecondName string `json:"second_name" binding:"required"`
	ThirdName  string `json:"third_name"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
}

type UpdateDTO struct {
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	ThirdName  string `json:"third_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RoleID     int    `json:"role_id"`
}
