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
var admins []m.User
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

	// Migrate the User, Announcements, Losts model to the db
	db.AutoMigrate(&m.User{}, &m.Announcement{}, &m.Lost{}, &m.Found{})

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
			IsAdmin:   false,
		},
		{
			Username:  "akhil_m",
			Password:  "MAC",
			FirstName: "Akhil",
			LastName:  "Maddukuri",
			Phone:     "09876543221",
			AptNo:     "HH-331",
			IsAdmin:   false,
		},
		{
			Username:  "bhavana_ch",
			Password:  "CHB",
			FirstName: "Bhavana",
			LastName:  "Chinthalapally",
			Phone:     "6789012345",
			AptNo:     "DD-447",
			IsAdmin:   false,
		},
		{
			Username:  "test_123",
			Password:  "123",
			FirstName: "Test",
			LastName:  "User",
			Phone:     "1234509876",
			AptNo:     "Apt123",
			IsAdmin:   false,
		},
		{
			Username:  "harshitha_m",
			Password:  "GCDC",
			FirstName: "Harshitha",
			LastName:  "Myadam",
			Phone:     "67878123345",
			AptNo:     "DD-233",
			IsAdmin:   false,
		},
		{
			Username:  "tejarocks",
			Password:  "VRA",
			FirstName: "Tejasri",
			LastName:  "Dontham",
			Phone:     "6783201145",
			AptNo:     "DD-244",
			IsAdmin:   false,
		},
		{
			Username:  "testuser3",
			Password:  "random",
			FirstName: "Test-3",
			LastName:  "User",
			Phone:     "9162201145",
			AptNo:     "X-123",
			IsAdmin:   false,
		},
		{
			Username:  "testuser4",
			Password:  "alsorandom",
			FirstName: "Teja",
			LastName:  "D",
			Phone:     "1234567891",
			AptNo:     "AA-123",
			IsAdmin:   false,
		},
		{
			Username:  "testuser5",
			Password:  "hehe",
			FirstName: "TestUser",
			LastName:  "Five",
			Phone:     "9848586878",
			AptNo:     "YY-174",
			IsAdmin:   false,
		},
		{
			Username:  "manish",
			Password:  "villages",
			FirstName: "Manish",
			LastName:  "Alluri",
			Phone:     "7656463626",
			AptNo:     "X-14",
			IsAdmin:   false,
		},
		{
			Username:  "anvitha",
			Password:  "dontknow",
			FirstName: "anvitha",
			LastName:  "choday",
			Phone:     "9738954545",
			AptNo:     "HH-331",
			IsAdmin:   false,
		},
		{
			Username:  "squareroot",
			Password:  "forgot",
			FirstName: "imaginary",
			LastName:  "M",
			Phone:     "6837773773",
			AptNo:     "GG-333",
			IsAdmin:   false,
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
		{
			Announcement_Id:       6,
			Admin_Id:              1,
			Announcement_Category: "Test",
			Announcement_Title:    "Test data",
			Venue:                 "Test data",
			Event_Description:     "Test data",
			CreatedAt:             db.NowFunc().UTC().Local(),
			UpdatedAt:             db.NowFunc().UTC().Local(),
		},
	}
	db.Create(&announcements)
	admins = []m.User{
		{
			Username:  "nitya_v2",
			Password:  "MNV",
			FirstName: "Nitya Vaishnavi",
			LastName:  "Mamillapalle",
			Phone:     "1234567890",
			IsAdmin:   true,
		},
		{
			Username:  "akhil_m2",
			Password:  "MAC",
			FirstName: "Akhil",
			LastName:  "Maddukuri",
			Phone:     "09876543221",
			IsAdmin:   true,
		},
		{
			Username:  "bhavana_ch2",
			Password:  "CHB",
			FirstName: "Bhavana",
			LastName:  "Chinthalapally",
			Phone:     "6789012345",
			IsAdmin:   true,
		},
		{
			Username:  "test_1234",
			Password:  "123",
			FirstName: "Test",
			LastName:  "User",
			Phone:     "1234509876",
			IsAdmin:   true,
		},
		{
			Username:  "harshitha_m2",
			Password:  "GCDC",
			FirstName: "Harshitha",
			LastName:  "Myadam",
			Phone:     "67878123345",
			IsAdmin:   true,
		},
	}
	db.Create(&admins)

	lostitems := []m.Lost{
		{
			LostType:    "Jewellery",
			Description: "lost near pool",
			ImagePath:   "https://image",
		},
		{
			UserId:      1,
			LostType:    "Bag",
			Description: "Contains my wallet and laptop",
			ImagePath:   "https://image2",
		},
		{
			LostType:    "Key",
			Description: "I don't have a spare key, Please return if found",
			ImagePath:   "https://image4",
		},
		{
			LostType:    "Heart",
			Description: "don't wanna be heartless",
			ImagePath:   "https://image5",
		},
	}
	db.Create(&lostitems)

	founditems := []m.Found{
		{
			FoundType:   "Jewellery",
			Description: "found near pool",
			ImagePath:   "https://image",
		},
		{
			UserId:      1,
			FoundType:   "Bag",
			Description: "Contains my wallet and laptop",
			ImagePath:   "https://image2",
		},
		{
			FoundType:   "Key",
			Description: "key found",
			ImagePath:   "https://image4",
		},
		{
			FoundType:   "Heart",
			Description: "i don't wanna return",
			ImagePath:   "https://image5",
		},
	}
	db.Create(&founditems)

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

func TestAdminCreateAnnouncementPassCase(t *testing.T) {
	login := m.Login{
		Username: "testadmin",
		Password: "TestAdmin@123",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/login", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		announcements := m.Announcement{
			Announcement_Id:       1,
			Admin_Id:              1,
			Announcement_Category: "Event Announcement",
			Announcement_Title:    "New Year Celebrations",
			Venue:                 "Stoneridge Leasing Office",
			Event_Description:     "We are celebrating New Year's!",
		}
		nr.Flush()
		body, _ := json.Marshal(announcements)
		req, _ := http.NewRequest("POST", "/createAnnouncement", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		req1.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr, req)
		assert.Equal(t, 200, nr.Code)
	}
}

func TestAdminEditAnnouncementPassCase(t *testing.T) {
	login := m.Login{
		Username: "testadmin",
		Password: "TestAdmin@123",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/login", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		announcement := m.Announcement{
			Announcement_Id:       1,
			Admin_Id:              1,
			Announcement_Category: "Event Announcement",
			Announcement_Title:    "New Year Celebrations",
			Venue:                 "Stoneridge Pool",
			Event_Description:     "We are celebrating New Year's!",
		}
		nr.Flush()
		body, _ := json.Marshal(announcement)
		req, _ := http.NewRequest("PUT", "/editAnnouncement", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		req1.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr, req)
		assert.Equal(t, 200, nr.Code)
	}
}

func TestAdminCreateAnnoncementFailCase(t *testing.T) {
	login := m.Login{
		Username: "testadmin",
		Password: "TestAdmin@123",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/login", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	if nr.Code == 200 {
		announcement := m.Announcement{
			Announcement_Id:       1,
			Admin_Id:              1,
			Announcement_Category: "Event Announcement",
			Announcement_Title:    "New Year Celebrations",
			Venue:                 "Stoneridge Leasing Office",
			Event_Description:     "We are celebrating New Year's!",
		}
		nr.Flush()
		w := httptest.NewRecorder()
		body, _ := json.Marshal(announcement)
		req, _ := http.NewRequest("POST", "/createAnnouncement", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		req1.Header.Set("credentials", "include")
		router.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
	}
}

func TestAdminDeleteAnnouncementFailCase(t *testing.T) {
	login := m.Login{
		Username: "testadmin",
		Password: "TestAdmin@123",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/login", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/deleteAnnouncement/-1", nil)
		req1.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(w, req)
		assert.Equal(t, 404, w.Code)

	}
}
