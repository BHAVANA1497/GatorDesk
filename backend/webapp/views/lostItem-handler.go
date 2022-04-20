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

//Posting a lost item into the Lost table. The user's id who has posted the lost item shall be retrieved from the sesssion
func PostLostItemView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		session := sessions.Default(c)
		var json l.Lost
		// try to bind the request json to the lost struct
		if err := c.ShouldBindJSON(&json); err != nil {
			// return bad request if field names are wrong
			// and if fields are missing
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//check if any user logged in
		if session.Get("uId") != nil {

			// strips HTML input from user for security purpose
			p := bluemonday.StripTagsPolicy()
			v := session.Get("uId")

			json.UserId = int64(v.(uint))
			json.LostType = p.Sanitize(json.LostType)
			json.Description = p.Sanitize(json.Description)
			json.ImagePath = p.Sanitize(json.ImagePath)
			json.IsFound = false

			// create the announcement
			result := db.Create(&json)

			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"result": "Lost item uploaded successfully!",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{

				"result": "User not loggedin",
			})
			return
		}

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

//API to list all the lost items to the Admin
func ListAllLostItemsView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		session := sessions.Default(c)

		// m, ok := session.(interface{})
		// if !ok {
		// 	return fmt.Errorf("want type map[string]interface{};  got %T", session)
		// }
		// for k, v := range m {
		// 	fmt.Println(k, "=>", v)
		// }

		if session.Get("uId") != nil {

			v := session.Get("uId")

			id := int64(v.(uint))

			var admin l.User
			db.Find(&admin, "id = ?", id)
			if admin.IsAdmin == true {
				var res []l.Lost
				db.Find(&res)
				//fmt.Print("shouldn't be here")
				c.JSON(http.StatusOK, gin.H{"data": res})
			} else {
				//fmt.Print("here")
				c.JSON(http.StatusUnauthorized, gin.H{

					"result": "login as admin to view this page",
				})
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{

				"result": "User not loggedin",
			})
			return
		}

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

func GetLostItemById(db *gorm.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		session := sessions.Default(c)

		if session.Get("uId") != nil {

			lostId, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			var res []l.Lost
			db.Find(&res, "id = ?", lostId)

			c.JSON(http.StatusOK, gin.H{"data": res})

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{

				"result": "User not loggedin",
			})
			return
		}
	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

//API to get all the lost items posted by a particular user (with a unique userId)
func GetAllLostItemsByUserId(db *gorm.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		session := sessions.Default(c)

		if session.Get("uId") != nil {

			v := session.Get("uId")

			id := int64(v.(uint))

			var lostitems []l.Lost
			db.Find(&lostitems, "user_id= ?", id)

			c.JSON(http.StatusOK, gin.H{"data": lostitems})
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{

				"result": "User not loggedin",
			})
			return
		}

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}

func LinkLostFoundView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json l.Lost
		// try to bind the request json to the Announcement struct
		if err := c.ShouldBindJSON(&json); err != nil {
			// return bad request if field names are wrong
			// and if fields are missing
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// strips HTML input from user for security purpose
		p := bluemonday.StripTagsPolicy()

		json.LostType = p.Sanitize(json.LostType)
		json.Description = p.Sanitize(json.Description)
		json.ImagePath = p.Sanitize(json.ImagePath)

		// get the existing post
		var lost l.Lost
		db.Find(&lost, "id = ?", json.ID)

		result := db.Model(&lost).Where("id=?", json.ID).Updates(json)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, lost)

	}

	// return the loginHandlerfunction
	return gin.HandlerFunc(fn)
}
