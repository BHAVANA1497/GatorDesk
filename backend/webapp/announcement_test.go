package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	m "webapp/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbName string = "test.db"
var storeName string = "testsecret"
var sessionName string = "testsession"
var users []m.User
var announcements []m.Announcement
var router *gin.Engine

func setupTestDb(dbName string) *gorm.DB {
	// Connection to the database with default configuration
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	// drop tables if exist
	db.Migrator().DropTable(&m.User{})
	db.Migrator().DropTable(&m.Announcement{})

	// Migrate the User & Product model to the db
	db.AutoMigrate(&m.User{}, &m.Announcement{})

	return db
}

func initData(db *gorm.DB) {
	users = []m.User{
		{
			Username:  "nitya_v",
			Password:  "MNV",
			FirstName: "Nitya Vaishnavi",
			LastName:  "Mamillapalle",
			Phone:     "1234567890",
			AptNo:     "GG-333",
		},
		{
			Username:  "akhil_m",
			Password:  "MAC",
			FirstName: "Akhil",
			LastName:  "Maddukuri",
			Phone:     "09876543221",
			AptNo:     "HH-331",
		},
		{
			Username:  "bhavana_ch",
			Password:  "CHB",
			FirstName: "Bhavana",
			LastName:  "Chinthalapally",
			Phone:     "6789012345",
			AptNo:     "DD-447",
		},
	}
	db.Create(&users)

	announcements = []m.Announcement{
		{
			Announcement_Id:       1,
			Admin_Id:              1,
			Announcement_Category: "Event Announcement",
			Announcement_Title:    "New Year Celebrations",
			Venue:                 "Stoneridge Leasing Office",
			Event_Description:     "We are celebrating New Year's!",
			CreatedAt:             db.NowFunc().UTC().Local(),
			UpdatedAt:             db.NowFunc().UTC().Local(),
		},
		{
			Announcement_Id:       2,
			Admin_Id:              1,
			Announcement_Category: "Maintainence Announcement",
			Announcement_Title:    "Fire Extinguisher repair",
			Venue:                 "At your own apartment",
			Event_Description:     "We are repairing your Fire extinguisher today at 6:00 pm",
			CreatedAt:             db.NowFunc().UTC().Local(),
			UpdatedAt:             db.NowFunc().UTC().Local(),
		},
		{
			Announcement_Id:       3,
			Admin_Id:              2,
			Announcement_Category: "Celebration",
			Announcement_Title:    "Mardi Grass",
			Venue:                 "Stoneridge Celebration Corner",
			Event_Description:     "We are celebrating Mardi Grass Lunch",
			CreatedAt:             db.NowFunc().UTC().Local(),
			UpdatedAt:             db.NowFunc().UTC().Local(),
		},
		{
			Announcement_Id:       4,
			Admin_Id:              2,
			Announcement_Category: "Celebration",
			Announcement_Title:    "Pool Party",
			Venue:                 "Stoneridge Pool",
			Event_Description:     "We are celebrating Stoneridge's one year ceremony",
			CreatedAt:             db.NowFunc().UTC().Local(),
			UpdatedAt:             db.NowFunc().UTC().Local(),
		},
		{
			Announcement_Id:       5,
			Admin_Id:              1,
			Announcement_Category: "Emergency Drill",
			Announcement_Title:    "Earthquare Disaster Emergency Drill",
			Venue:                 "Stoneridge Playground",
			Event_Description:     "We are training you for disasters.",
			CreatedAt:             db.NowFunc().UTC().Local(),
			UpdatedAt:             db.NowFunc().UTC().Local(),
		},
	}
	db.Create(&announcements)
}

func TestMain(m *testing.M) {
	// setup database
	db := setupTestDb(dbName)

	// init data
	initData(db)

	// setup router
	router = SetupRouter(db, storeName, sessionName)

	code := m.Run()

	os.Exit(code)
}

func TestListAllAnnouncementsView(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/listAllAnnouncements", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	b, _ := json.Marshal(announcements)
	res := strings.TrimLeft(w.Body.String(), "{\"data\":")
	res1 := strings.TrimRight(res, "}")
	assert.Equal(t, string(b), res1)

}
