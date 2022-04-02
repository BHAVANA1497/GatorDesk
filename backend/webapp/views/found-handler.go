package views

import (
	"net/http"

	l "webapp/model"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)

//Posting a found item into the Lost table. The user's id who has posted the found item shall be retrieved from the sesssion
func PostFoundItemView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json l.Found
		// try to bind the request json to the Login struct
		if err := c.ShouldBindJSON(&json); err != nil {
			// return bad request if field names are wrong
			// and if fields are missing
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// strips HTML input from user for security purpose
		p := bluemonday.StripTagsPolicy()

		json.FoundType = p.Sanitize(json.FoundType)
		json.Description = p.Sanitize(json.Description)
		json.ImagePath = p.Sanitize(json.ImagePath)

		// create the announcement
		result := db.Create(&json)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "Found item has been updated successfully in the database!",
		})
	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

func ListAllFoundItemsView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var res []l.Found
		db.Find(&res)

		c.JSON(http.StatusOK, gin.H{"data": res})

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

func GetAllFoundItemsByUserId(db *gorm.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		var json l.Found
		// try to bind the request json to the Announcement struct
		if err := c.ShouldBindJSON(&json); err != nil {
			// return bad request if field names are wrong
			// and if fields are missing
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// strips HTML input from user for security purpose
		p := bluemonday.StripTagsPolicy()

		json.FoundType = p.Sanitize(json.FoundType)
		json.Description = p.Sanitize(json.Description)
		json.ImagePath = p.Sanitize(json.ImagePath)

		var res []l.Lost
		db.Find(&res, "UserId = ?", json.UserId)

		c.JSON(http.StatusOK, gin.H{"data": res})

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}
