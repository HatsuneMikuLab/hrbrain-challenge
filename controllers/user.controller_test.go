package controllers

import (
	"encoding/json"
	"bytes"
	"github.com/HatsuneMikuLab/hrbrain-challenge/db"
	"github.com/HatsuneMikuLab/hrbrain-challenge/services"
	"github.com/HatsuneMikuLab/hrbrain-challenge/modules"
	"net/http"
	"testing"
	"net/http/httptest"
)

// coverage is bad, I am just too lazy
func TestAddUser(t *testing.T) {
	rec := httptest.NewRecorder()
	db := db.Connect2db("postgres://hzvlouaa:Gy5nzrwrQnw8IH9emXtEdH-RNPdI0SAd@kashin.db.elephantsql.com/hzvlouaa")
	ctrl := NewUserController(
		services.NewUserService(db), 
		services.NewEvaluationService(),
		services.NewCacheService(),
	)
	body := json.Marshal(&User{ ID: "Ophelia", Email: "ophelia@gmail.com" })
	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Could not create request %s", err.Error())
	}
	handler := http.HandlerFunc(ctrl.GetUserByID)
	handler.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Response status should be 200, but got %v", res.StatusCode)
	} 
}