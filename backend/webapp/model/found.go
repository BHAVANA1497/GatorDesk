package model

import "gorm.io/gorm"

//Found table
type Found struct {
	gorm.Model
	UserId      int64  `gorm:"not_null"`
	FoundType   string `form:"found_type" json:"found_type"`
	Description string `form:"description" json:"description"`
	ImagePath   string `form:"image_path" json:"image_path"`
}
