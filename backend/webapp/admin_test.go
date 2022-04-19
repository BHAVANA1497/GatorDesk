package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	m "webapp/model"

	"github.com/stretchr/testify/assert"
)

func TestAddAdminPassCase(t *testing.T) {
	login := m.Login{
		Username: "testadmin",
		Password: "TestAdmin@123",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/adminlogin", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		admins := m.User{
			Username:  "Akhil",
			Password:  "password",
			FirstName: "Akhil",
			LastName:  "Maddukuri",
			Phone:     "+157355155",
		}
		nr.Flush()
		body, _ := json.Marshal(admins)
		req, _ := http.NewRequest("POST", "/addAdmin", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		req1.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr, req)
		assert.Equal(t, 200, nr.Code)
	}
}

func TestAddAdminFailCase(t *testing.T) {
	login := m.Login{
		Username: "testadmin",
		Password: "TestAdmin@123",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/adminlogin", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	if nr.Code == 200 {
		admin := m.User{
			Username:  "Akhil",
			Password:  "password",
			FirstName: "Akhil",
			LastName:  "Maddukuri",
			Phone:     "+157355155",
		}
		nr.Flush()
		w := httptest.NewRecorder()
		body, _ := json.Marshal(admin)
		req, _ := http.NewRequest("POST", "/createAnnouncement", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		req1.Header.Set("credentials", "include")
		router.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
	}
}
