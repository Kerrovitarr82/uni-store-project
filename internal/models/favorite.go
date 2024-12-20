package models

import "time"

type Favorite struct {
	ID         int  `json:"id"`
	CustomerID int  `json:"customer_id"`
	User       User `gorm:"foreignKey:CustomerID" json:"user"`
}

type ProductInFavorite struct {
	ProductID  int       `json:"product_id"`
	Product    Product   `gorm:"foreignKey:ProductID" json:"product"`
	FavoriteID int       `json:"favorite_id"`
	Favorite   Favorite  `gorm:"foreignKey:FavoriteID" json:"favorite"`
	CreatedAt  time.Time `json:"created_at"`
}
