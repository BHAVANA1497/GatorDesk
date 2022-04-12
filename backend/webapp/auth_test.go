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

func TestRegisterPassCase(t *testing.T) {
	user_details := m.User{
		Username:  "Vamsee",
		Password:  "nothing",
		FirstName: "Vamsee",
		LastName:  "Raj",
		Phone:     "8928954545",
		AptNo:     "B-1333",
	}
	body, err := json.Marshal(user_details)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/signUp", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestRegisterFailCase(t *testing.T) {
	user_details := m.User{

		Username:  "akhil_m",
		Password:  "MAC",
		FirstName: "Akhil",
		LastName:  "Maddukuri",
		Phone:     "09876543221",
		AptNo:     "HH-331",
	}
	body, err := json.Marshal(user_details)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/signUp", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

}
func TestDeleteUserPassCase(t *testing.T) {

	w := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/deleteUser/10", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

}

func TestDeleteUserFailCase(t *testing.T) {

	w := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/deleteUser/10", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

}

func TestAdminLoginPassCase(t *testing.T) {
	login := m.Login{Username: "nitya_v", Password: "MNV"}

	out, err := json.Marshal(login)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/adminlogin", strings.NewReader(string(out)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

}

func TestAdminLoginFailCase(t *testing.T) {
	login := m.Login{Username: "nitya_v", Password: "wrongpassword@123"}

	out, err := json.Marshal(login)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/adminlogin", strings.NewReader(string(out)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

}
