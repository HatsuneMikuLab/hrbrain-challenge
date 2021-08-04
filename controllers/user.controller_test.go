package controllers

import (
	"io/ioutil"
	"github.com/HatsuneMikuLab/hrbrain-challenge/models"
	"encoding/json"
	"bytes"
	"github.com/HatsuneMikuLab/hrbrain-challenge/db"
	"github.com/HatsuneMikuLab/hrbrain-challenge/services"
	"net/http"
	"testing"
	"net/http/httptest"
)

func prepare4test() *UsersController {
	db := db.Connect2db("") //put here your DB connection
	return NewUserController(
		services.NewUserService(db), 
		services.NewEvaluationService(),
		services.NewCacheService(),
	)
} 

// coverage is bad, I am just too lazy
func TestAddUser(t *testing.T) {
	user := &models.User{}
	expectedID := "Nezuko"
	expectedEmail := "nezuko@gmail.com"
	rec := httptest.NewRecorder()
	ctrl := prepare4test()
	
	body, _ := json.Marshal(&models.User{ ID: expectedID, Email: expectedEmail })
	req, err := http.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Could not create request %s", err.Error())
	}
	t.Logf("REQ: %v", req)
	handler := http.HandlerFunc(ctrl.AddUser)
	handler.ServeHTTP(rec, req)

	
	res := rec.Result()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read body responset %s", err.Error())
	}

	t.Logf("RES: %v", string(resBody))

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Response status should be 200, but got %v", res.StatusCode)
	} 
	if json.Unmarshal(resBody, user) != nil  {
		t.Fatalf("Response body should have a JSON format %v", resBody)
	}
	if user.ID != expectedID || user.Email != expectedEmail {
		t.Fatalf("Response body is expected to have ID: %s and Email %s, but got %v", expectedID, expectedEmail, resBody)
	}
}