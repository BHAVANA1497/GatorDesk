package model

import "gorm.io/gorm"

//Found table
type Found struct {
	gorm.Model
	FoundId     int64  `gorm:"primary_key;auto_increment;not_null;unique" form:"found_id" json:"found_id"`
	UserId      int64  `gorm:"not_null"`
	FoundType   string `form:"found_type" json:"found_type"`
	Description string `form:"description" json:"description"`
	ImagePath   string `form:"image_path" json:"image_path"`
}
