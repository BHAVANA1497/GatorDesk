package views

import (
	"net/http"

	m "webapp/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)

/*
	Add Admin - adds a new admin

	if admin already exists or unable to create returns StatusBadRequest
	else Status OK 200
*/
func AddAdminView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json m.Admin
		// try to bind the request json to the User struct
		if err := c.ShouldBindJSON(&json); err != nil {
			// return bad request if field names are wrong
			// and if fields are missing
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// strips HTML input from strings preventing XSS
		p := bluemonday.StripTagsPolicy()
		json.Adminname = p.Sanitize(json.Adminname)
		json.Password = p.Sanitize(json.Password)
		json.FirstName = p.Sanitize(json.FirstName)
		json.LastName = p.Sanitize(json.LastName)
		json.Phone = p.Sanitize(json.Phone)

		// check db if the admin exists
		var admin m.Admin
		db.Find(&admin, "adminname = ?", json.Adminname)
		// return error if the admin exists
		if admin != (m.Admin{}) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Admin with these credentials already exists!"})
			return
		}

		// create the user
		result := db.Create(&json)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "Succesfully added an admin!",
		})
	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

func AdminLoginView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		session := sessions.Default(c)
		var json m.AdminLogin
		// try to bind the request json to the Login struct
		if err := c.ShouldBindJSON(&json); err != nil {
			// return bad request if field names are wrong
			// and if fields are missing
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var admins []m.Admin

		// strips HTML input from strings preventing XSS
		p := bluemonday.StripTagsPolicy()
		adminname := p.Sanitize(json.Adminname)
		password := p.Sanitize(json.Password)

		// DB query to search for username and password and store the results in users
		db.Find(&admins, "adminname = ? AND password = ?", adminname, password)

		// if user found return success
		if len(admins) > 0 {
			session.Set("uId", admins[0].ID)
			session.Save()
			c.JSON(http.StatusOK, gin.H{
				"result": "Successful Login!",
			})
			return
		}

		// return unauthorized status
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "Invalid Admin name and/or Password",
		})
	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}
