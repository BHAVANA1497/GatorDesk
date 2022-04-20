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

func TestPostLostItemPassCase(t *testing.T) {
	login := m.Login{
		Username: "tejarocks",
		Password: "VRA",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/login", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		lostitem := m.Lost{
			LostType:    "Jewellery",
			Description: "lost near pool",
			ImagePath:   "https://image",
		}
		nr.Flush()
		body, _ := json.Marshal(lostitem)
		req, _ := http.NewRequest("POST", "/postLostItem", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		req1.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr, req)
		assert.Equal(t, 200, nr.Code)
	}

}

func TestPostLostItemFailCase(t *testing.T) {
	login := m.Login{
		Username: "tejarockss",
		Password: "na",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/login", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		lostitem := m.Lost{
			LostType:    "Jewellery",
			Description: "lost near pool",
			ImagePath:   "https://image",
		}
		nr.Flush()
		body, _ := json.Marshal(lostitem)
		req, _ := http.NewRequest("POST", "/postLostItem", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr, req)
		assert.Equal(t, 401, nr.Code)
	} else {
		assert.Equal(t, 401, nr.Code)
	}

}

func TestListLostItemsPassCase(t *testing.T) {
	login := m.Login{
		Username: "nitya_v2",
		Password: "MNV",
	}
	payload, _ := json.Marshal(login)
	nr := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/adminlogin", strings.NewReader(string(payload)))
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("credentials", "include")
	router.ServeHTTP(nr, req1)
	cookieValue := nr.Result().Header.Get("Set-Cookie")
	if nr.Code == 200 {
		//fmt.Print("here")
		nr.Flush()
		req, _ := http.NewRequest("GET", "/listAllLostItems", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr, req)
		assert.Equal(t, 200, nr.Code)
	} else {
		assert.Equal(t, 200, nr.Code)
	}

}

func TestListLostItemsFailCase(t *testing.T) {
	login := m.Login{
		Username: "tejarocks",
		Password: "VRA",
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

		req, _ := http.NewRequest("GET", "/listAllLostItems", nil)
		req.Header.Set("Content-Type", "application/json")
		req1.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr2, req)
		//	fmt.Print(nr2.Code)
		assert.Equal(t, 401, nr2.Code)
	} else {
		assert.Equal(t, 401, nr.Code)
	}

}

func TestListLostItemsByUserPassCase(t *testing.T) {
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
		req, _ := http.NewRequest("GET", "/listAllLostItemsByUserId", nil)
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

func TestListLostItemsByUserFailCase(t *testing.T) {

	nr := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/listAllLostItemsByUserId", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("credentials", "include")
	req.Header.Set("Cookie", "")
	router.ServeHTTP(nr, req)

	assert.Equal(t, 401, nr.Code)

}
