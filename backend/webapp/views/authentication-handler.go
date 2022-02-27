package views

import (
	"net/http"
	m "webapp/model"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)

/*
	Sign Up View - creates a new user

	if user exists or unable to create returns StatusBadRequest
	else Status OK 200
*/
func SignUpView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json m.User
		// try to bind the request json to the User struct
		if err := c.ShouldBindJSON(&json); err != nil {
			// return bad request if field names are wrong
			// and if fields are missing
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// strips HTML input from strings preventing XSS
		p := bluemonday.StripTagsPolicy()
		json.Username = p.Sanitize(json.Username)
		json.Password = p.Sanitize(json.Password)
		json.FirstName = p.Sanitize(json.FirstName)
		json.LastName = p.Sanitize(json.LastName)
		json.Phone = p.Sanitize(json.Phone)
		json.AptNo = p.Sanitize(json.AptNo)

		// check db if the username exists
		var user m.User
		db.Find(&user, "username = ?", json.Username)
		// return error if the user exists
		if user != (m.User{}) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username Already Exists!"})
			return
		}

		// create the user
		result := db.Create(&json)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "Succesfully signed up the user!",
		})
	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}
