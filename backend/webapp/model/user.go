package model

import "gorm.io/gorm"

// User table
type User struct {
	gorm.Model
	Username  string `gorm:"primary_key;not_null;unique" form:"username" json:"username" binding:"required"`
	Password  string `gorm:"not_null" form:"password" json:"password" binding:"required"`
	FirstName string `gorm:"not_null" form:"firstname" json:"firstname" binding:"required"`
	LastName  string `gorm:"not_null" form:"lastname" json:"lastname" binding:"required"`
	Phone     string `gorm:"not_null" form:"phone" json:"phone" binding:"required"`
	AptNo     string `gorm:"not_null" form:"aptNo" json:"aptNo" binding:"required"`
}

// Login struct, mapping json to the variables
type Login struct {
	Username string `gorm:"not_null" form:"username" json:"username" binding:"required"`
	Password string `gorm:"not_null" form:"password" json:"password" binding:"required"`
}
