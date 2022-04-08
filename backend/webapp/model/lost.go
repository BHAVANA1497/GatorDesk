package model

import "gorm.io/gorm"

//Lost table
type Lost struct {
	gorm.Model
	UserId      int64  `gorm:"not_null"`
	LostType    string `form:"lost_type" json:"lost_type"`
	Description string `form:"description" json:"description"`
	ImagePath   string `json:"image_path"`
	IsFound     bool   `json:"is_found"`
	FoundId     int64  `json:"found_id"`
}
