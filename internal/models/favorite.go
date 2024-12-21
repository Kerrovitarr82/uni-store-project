package models

type Favorite struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
	Games  []Game `gorm:"many2many:favorite_games;" json:"games"`
}
