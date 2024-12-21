package models

type ShoppingCart struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	UserId int    `json:"user_id"`
	User   User   `gorm:"foreignKey:UserId" json:"user"`
	Games  []Game `gorm:"many2many:games_in_cart;" json:"games"`
}
