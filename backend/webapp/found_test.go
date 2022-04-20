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
