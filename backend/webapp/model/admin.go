package model

import "gorm.io/gorm"

//Admin table
type Admin struct {
	gorm.Model
	IsAdmin   bool
	Adminname string `gorm:"not_null" form:"adminname" json:"adminname" binding:"required"`
	Password  string `gorm:"not_null" form:"password" json:"password" binding:"required"`
	FirstName string `gorm:"not_null" form:"firstname" json:"firstname" binding:"required"`
	LastName  string `gorm:"not_null" form:"lastname" json:"lastname" binding:"required"`
	Phone     string `gorm:"not_null" form:"phone" json:"phone" binding:"required"`
}

// Login struct, mapping json to the variables
type AdminLogin struct {
	Adminname string `gorm:"not_null" form:"adminname" json:"adminname" binding:"required"`
	Password  string `gorm:"not_null" form:"password" json:"password" binding:"required"`
}
