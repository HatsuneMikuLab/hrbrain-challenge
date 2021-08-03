package controllers

import (
	"github.com/HatsuneMikuLab/hrbrain-challenge/services"
	"encoding/json"
	"net/http"
)

type UsersController struct {
	UsersService services.IUsersService
}

func GetUserByID(res http.ResponseWriter, req *http.Request) {
	user := &User{}
	
	json.NewEncoder(res).Encode(user)
}