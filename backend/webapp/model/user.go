package model

import "gorm.io/gorm"

// User table
type User struct {
	gorm.Model
	IsAdmin   bool
	Username  string `gorm:"not_null" form:"username" json:"username" binding:"required"`
	Password  string `gorm:"not_null" form:"password" json:"password" binding:"required"`
	FirstName string `gorm:"not_null" form:"firstname" json:"firstname" binding:"required"`
	LastName  string `gorm:"not_null" form:"lastname" json:"lastname"`
	Phone     string `gorm:"not_null" form:"phone" json:"phone"`
	AptNo     string `gorm:"not_null" form:"aptNo" json:"aptNo"`
}

// Login struct, mapping json to the variables
type Login struct {
	Username string `gorm:"not_null" form:"username" json:"username" binding:"required"`
	Password string `gorm:"not_null" form:"password" json:"password" binding:"required"`
}
