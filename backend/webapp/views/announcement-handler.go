package views

import (
	"net/http"
	l "webapp/model"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)

// Announcements View - wrapper func around the announcements handler method
// this is because we have passed the db variable
// could've passed db variable as middleware but it is strictly against doing it in go documentation
func CreateAnnouncementView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json l.Announcement
		// try to bind the request json to the Login struct
		if err := c.ShouldBindJSON(&json); err != nil {
			// return bad request if field names are wrong
			// and if fields are missing
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// strips HTML input from user for security purpose
		p := bluemonday.StripTagsPolicy()

		//json.Announcement_Id = p.Sanitize(json.Announcement_Id)
		//json.Admin_Id = p.Sanitize(json.Admin_Id)
		json.Announcement_Category = p.Sanitize(json.Announcement_Category)
		json.Announcement_Title = p.Sanitize(json.Announcement_Title)
		json.Venue = p.Sanitize(json.Venue)
		json.Event_Description = p.Sanitize(json.Event_Description)
		//json.CreatedAt = p.Sanitize(json.CreatedAt)
		//json.UpdatedAt = p.Sanitize(json.UpdatedAt)

		// create the announcement
		result := db.Create(&json)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "Announcement created successfully",
		})
	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

// Announcements View - wrapper func around the announcements handler method
// this is because we have passed the db variable
// could've passed db variable as middleware but it is strictly against doing it in go documentation
func EditAnnouncementView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json l.Announcement
		// try to bind the request json to the Announcement struct
		if err := c.ShouldBindJSON(&json); err != nil {
			// return bad request if field names are wrong
			// and if fields are missing
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// strips HTML input from user for security purpose
		p := bluemonday.StripTagsPolicy()

		//json.Announcement_Id = p.Sanitize(json.Announcement_Id)
		//json.Admin_Id = p.Sanitize(json.Admin_Id)
		json.Announcement_Category = p.Sanitize(json.Announcement_Category)
		json.Announcement_Title = p.Sanitize(json.Announcement_Title)
		json.Venue = p.Sanitize(json.Venue)
		json.Event_Description = p.Sanitize(json.Event_Description)
		//json.CreatedAt = p.Sanitize(json.CreatedAt)
		//json.UpdatedAt = p.Sanitize(json.UpdatedAt)

		// get the existing post
		var ann l.Announcement
		db.Find(&ann, "Announcement_Id = ?", json.Announcement_Id)

		result := db.Model(&ann).Where("Announcement_Id=?", json.Announcement_Id).Updates(json)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, ann)

		/*
			// return error if the user does not exist
			if ann.Announcement_Id == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Announcement does not exist!"})
				return
			}
		*/

		/* c.JSON(http.StatusOK, gin.H{
			"result": "Announcement edited successfully",
		}) */

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}
