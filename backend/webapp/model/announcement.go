package model

import (
	//"gorm.io/gorm"
	"time"

)

type Announcement struct {
	Announcement_Id int64 `gorm:"primary_key;auto_increment;not_null;unique" form:"announcementId" json:"announcementId"` 
	Admin_Id int64 `gorm:"not_null" form:"adminId" json:"adminId" binding:"required"`
	Announcement_Category  string `form:"announcementCategory" json:"announcementCategory" binding:"required"`
	Announcement_Title string `gorm:"not_null" form:"announcementTitle" json:"announcementTitle" binding:"required"` 
	Venue string `form:"venue" json:"venue" binding:"required"`
	Event_Description string `form:"eventDescription" json:"eventDescription" binding:"required"`
	CreatedAt time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `form:"updatedAt" json:"updatedAt"`
}