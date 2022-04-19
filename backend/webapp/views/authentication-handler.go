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
		json.IsAdmin = false
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

		c.JSON(http.StatusOK, gin.H{"result": "SingUp Successfully!", "username": json.Username})

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

/* Login View  - checks db for the user and password
   -if the user doesn't exists return UnAuthorized status code
    else return succcess code
   -wrapper func around the main login handler method
    this is because we have passed the db variable
    could've passed db variable as middleware but it is strictly against doing it in go documentation
*/
func LoginView(db *gorm.DB) gin.HandlerFunc {
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
		var users []m.User

		// strips HTML input from strings preventing XSS
		p := bluemonday.StripTagsPolicy()
		username := p.Sanitize(json.Username)
		password := p.Sanitize(json.Password)

		// DB query to search for username and password and store the results in users
		db.Find(&users, "username = ? AND password = ?", username, password)

		// if user found return success
		if len(users) > 0 {
			session.Set("uId", users[0].ID)
			session.Save()
			c.JSON(http.StatusOK, gin.H{
				"result":   "Successful Login!",
				"username": users[0].Username,
				"IsAdmin":  users[0].IsAdmin,
			})
			return
		}

		// return unauthorized status
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "Invalid User name and/or Password",
		})
	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		userId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var rec m.User
		if err := db.Where("id = ?", userId).First(&rec).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		db.Delete(&rec)
		c.JSON(http.StatusOK, gin.H{"data": true})

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

// Logout view - to remove the user from the session
func LogoutView(c *gin.Context) {
	session := sessions.Default(c)

	v := session.Get("uId")
	if v == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "User not logged in",
		})
		return
	}

	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()

	v = session.Get("uId")
	if v == nil {
		c.JSON(http.StatusOK, gin.H{
			"result": "Logout successful",
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"result": "Logout failed",
	})
}
