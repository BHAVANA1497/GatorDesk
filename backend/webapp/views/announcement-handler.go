package views

import (
	"net/http"
	"strconv"
	l "webapp/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)

// Announcements View - wrapper func around the announcements handler method
// this is because we have passed the db variable
// could've passed db variable as middleware but it is strictly against doing it in go documentation
func CreateAnnouncementView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		session := sessions.Default(c)
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

		json.Announcement_Category = p.Sanitize(json.Announcement_Category)
		json.Announcement_Title = p.Sanitize(json.Announcement_Title)
		json.Venue = p.Sanitize(json.Venue)
		json.Event_Description = p.Sanitize(json.Event_Description)

		//check if any user is logged-in
		if session != nil {

			v := session.Get("uId")

			var admin l.Admin
			db.Find(&admin, "id = ?", v)

			//check if logged-in user is admin
			if v == admin.ID {

				json.Admin_Id = int64(admin.ID)
				// create the announcement
				result := db.Create(&json)

				if result.Error != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"result": "Announcement created successfully",
				})
			} else {
				//fmt.Print(" debug print 1- detected1\n")
				c.JSON(http.StatusUnauthorized, gin.H{

					"result": "Login in Admin mode to create an announcement",
				})
			}

		} else {
			//fmt.Print(" debug print 2- detected2\n")
			c.JSON(http.StatusUnauthorized, gin.H{

				"result": "Login in Admin mode to create an announcement",
			})

		}
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

		json.Announcement_Category = p.Sanitize(json.Announcement_Category)
		json.Announcement_Title = p.Sanitize(json.Announcement_Title)
		json.Venue = p.Sanitize(json.Venue)
		json.Event_Description = p.Sanitize(json.Event_Description)

		// get the existing post
		var ann l.Announcement
		db.Find(&ann, "Announcement_Id = ?", json.Announcement_Id)

		result := db.Model(&ann).Where("Announcement_Id=?", json.Announcement_Id).Updates(json)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, ann)

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

func ListAnnouncementsView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var res []l.Announcement
		db.Find(&res)

		c.JSON(http.StatusOK, gin.H{"data": res})

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

func DeleteAnnouncementView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		announcementId, err := strconv.Atoi(c.Param("announcementId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var rec l.Announcement
		if err := db.Where("Announcement_Id = ?", announcementId).First(&rec).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		db.Delete(&rec)
		c.JSON(http.StatusOK, gin.H{"data": true})

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}
