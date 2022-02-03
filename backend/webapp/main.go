package main

// Imports Gin, Gorm, models and views
import (
	m "webapp/model"
	a "webapp/views"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Driver function which starts the server
func main() {

	// Connection to the database with default configuration
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the Announcement model to the db
	db.AutoMigrate(&m.Announcement{})

	// setting up the webserver with default config
	r := gin.New()

	// Adding logger to the middleware
	//r.Use(gin.Logger())

	// Using default recovery mechanism in case of any unexpected crashes in webserver
	//r.Use(gin.Recovery())

	// **** END POINTS ****
	//r.GET("/listAllAnouncements", a.ListAnnouncementsView(db))
	//r.POST("/editAnnoucements/{id}", a.EditAnnouncementView(db))
	r.POST("/createAnnoucement", a.CreateAnnouncementView(db))
	//r.POST("/deleteAnnoucement/{id}", a.DeleteAnnouncementView(db))

	// starts server and listens on port 8080
	r.Run(":8181")
}
