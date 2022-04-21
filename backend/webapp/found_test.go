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

func TestPostFoundItemPassCase(t *testing.T) {
	login := m.Login{
		Username: "tejarocks",
		Password: "VRA",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	nr2 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/login", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		founditem := m.Found{
			FoundType:   "Jewellery",
			Description: "Found at pool",
			ImagePath:   "https://image",
		}
		nr.Flush()
		body, _ := json.Marshal(founditem)
		req, _ := http.NewRequest("POST", "/postFoundItem", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr2, req)
		assert.Equal(t, 200, nr2.Code)
	} else {
		assert.Equal(t, 200, nr.Code)
	}

}

func TestPostFoundItemFailCase(t *testing.T) {
	login := m.Login{
		Username: "tejarocks",
		Password: "na",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	nr2 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/login", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		founditem := m.Found{
			FoundType:   "Jewellery",
			Description: "Found at pool",
			ImagePath:   "https://image",
		}
		nr.Flush()
		body, _ := json.Marshal(founditem)
		req, _ := http.NewRequest("POST", "/postFoundItem", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr2, req)
		assert.Equal(t, 401, nr2.Code)
	} else {
		assert.Equal(t, 401, nr.Code)
	}

}
func TestListFoundItemsPassCase(t *testing.T) {
	login := m.Login{
		Username: "nitya_v2",
		Password: "MNV",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	nr2 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/adminlogin", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		//fmt.Print("here")
		nr.Flush()
		req, _ := http.NewRequest("GET", "/listAllFoundItems", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr2, req)
		assert.Equal(t, 200, nr2.Code)
	} else {
		assert.Equal(t, 200, nr.Code)
	}

}

func TestListFoundItemsFailCase(t *testing.T) {
	login := m.Login{
		Username: "nitya_v2",
		Password: "wrongpassword",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	nr2 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/adminlogin", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		//fmt.Print("here")
		nr.Flush()
		req, _ := http.NewRequest("GET", "/listAllFoundItems", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr2, req)
		assert.Equal(t, 401, nr2.Code)
	} else {
		assert.Equal(t, 401, nr.Code)
	}

}

func TestListFoundItemsByUserPassCase(t *testing.T) {
	login := m.Login{
		Username: "nitya_v",
		Password: "MNV",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	nr2 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/adminlogin", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		//fmt.Print("here")
		nr.Flush()
		req, _ := http.NewRequest("GET", "/listAllFoundItemsByUserId", nil)
		req.Header.Set("Content-Type", "application/json")
		req1.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr2, req)
		//fmt.Print(nr2.Body)
		assert.Equal(t, 200, nr2.Code)
	} else {
		assert.Equal(t, 200, nr.Code)
	}

}
func TestListFoundItemsByUserFailCase(t *testing.T) {

	nr := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/listAllFoundItemsByUserId", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("credentials", "include")
	req.Header.Set("Cookie", "")
	router.ServeHTTP(nr, req)

	assert.Equal(t, 401, nr.Code)

}
