package views

import (
	"net/http"
	"strconv"

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
		json.IsAdmin = true
		json.Username = p.Sanitize(json.Username)
		json.Password = p.Sanitize(json.Password)
		json.FirstName = p.Sanitize(json.FirstName)
		json.LastName = p.Sanitize(json.LastName)
		json.Phone = p.Sanitize(json.Phone)
		//** handle apt num case here **

		// check db if the admin exists
		var admin m.User
		db.Find(&admin, "username = ?", json.Username)
		// return error if the admin exists
		if admin != (m.User{}) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User with these credentials already exists!"})
			return
		}

		// create the admin
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
		var json m.Login
		// try to bind the request json to the Login struct
		if err := c.ShouldBindJSON(&json); err != nil {
			// return bad request if field names are wrong
			// and if fields are missing
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var admins []m.User

		// strips HTML input from strings preventing XSS
		p := bluemonday.StripTagsPolicy()
		adminname := p.Sanitize(json.Username)
		password := p.Sanitize(json.Password)

		// DB query to search for username and password and store the results in users
		db.Find(&admins, "username = ? AND password = ?", adminname, password)

		// if user found return success
		if len(admins) > 0 {
			v := session.Get("uId")
			if v == nil {
				session.Set("uId", admins[0].ID)
				session.Save()
				c.JSON(http.StatusOK, gin.H{
					"result":   "Successful Login!",
					"username": admins[0].Username,
					"isAdmin":  admins[0].IsAdmin,
				})
				return

			} else if v == admins[0].ID {

				c.JSON(http.StatusBadRequest, gin.H{
					"result": "Admin already loggedin",
				})
				return
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"result": "Another user loggedin, cannot login",
				})
				return

			}

		}

		// return unauthorized status
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "Invalid Admin name and/or Password",
		})
	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

func AdminDeleteView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		adminId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var rec m.User
		if err := db.Where("id = ?", adminId).First(&rec).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		db.Delete(&rec)
		c.JSON(http.StatusOK, gin.H{"data": true})

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

func ListAllAdminsView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var res []m.User
		db.Find(&res, "is_admin = ?", true)

		c.JSON(http.StatusOK, gin.H{"data": res})

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}
