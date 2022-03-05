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

func TestLoginPassCase(t *testing.T) {
	login := m.Login{Username: "nitya_v", Password: "MNV"}

	out, err := json.Marshal(login)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/login", strings.NewReader(string(out)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestLoginFailCase(t *testing.T) {
	login := m.Login{Username: "nitya_v", Password: "wrongpassword@123"}

	out, err := json.Marshal(login)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/login", strings.NewReader(string(out)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)

}
