package main

// Imports Gin, Gorm, models and views
import (
	m "webapp/model"
	a "webapp/views"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func SetupRouter(db *gorm.DB, storeName string, sessionName string) *gin.Engine {
	// setting up the webserver with default config
	r := gin.New()

	// Check every request if allowed for CORS
	r.Use(CORSMiddleware())

	// Adding logger to the middleware
	//r.Use(gin.Logger())

	// Using default recovery mechanism in case of any unexpected crashes in webserver
	r.Use(gin.Recovery())

	store := cookie.NewStore([]byte(storeName))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})
	r.Use(sessions.Sessions(sessionName, store))

	// **** END POINTS ****
	r.POST("/signUp", a.SignUpView(db))
	r.POST("/login", a.LoginView(db))
	r.DELETE("/deleteUser/:id", a.DeleteUser(db))
	r.POST("/logout", a.LogoutView)
	r.GET("/listAllAnnouncements", a.ListAnnouncementsView(db))
	r.PUT("/editAnnouncement", a.EditAnnouncementView(db))
	r.POST("/createAnnouncement", a.CreateAnnouncementView(db))
	r.DELETE("/deleteAnnouncement/:announcementId", a.DeleteAnnouncementView(db))
	r.POST("/adminlogin", a.AdminLoginView(db))
	r.POST("/addAdmin", a.AddAdminView(db))
	r.DELETE("/deleteAdmin/:id", a.AdminDeleteView(db))
	r.POST("/postLostItem", a.PostLostItemView(db))
	r.GET("/listAllLostItems", a.ListAllLostItemsView(db))
	r.GET("/listAllLostItemsByUserId", a.GetAllLostItemsByUserId(db))
	r.POST("/postFoundItem", a.PostFoundItemView(db))
	r.GET("/listAllFoundItems", a.ListAllFoundItemsView(db))

	return r
}

// Driver function which starts the server
func main() {

	// Connection to the database with default configuration
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the User model to the db
	db.AutoMigrate(&m.User{})

	// Migrate the Announcement model to the db
	db.AutoMigrate(&m.Announcement{})

	//migrate the admin model to the db
	db.AutoMigrate(&m.Admin{})

	//migrate the lost model to the db
	db.AutoMigrate(&m.Lost{})

	// setting up the webserver with default config
	storeName := "mainsecret"
	sessionName := "mainsession"
	r := SetupRouter(db, storeName, sessionName)

	// starts server and listens on port 8181
	r.Run(":8181")
}
