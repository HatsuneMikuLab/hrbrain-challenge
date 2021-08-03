package controllers

import (
	"fmt"
	"github.com/HatsuneMikuLab/hrbrain-challenge/models"
	"github.com/HatsuneMikuLab/hrbrain-challenge/services"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type UsersController struct {
	UsersService services.IUsersService
	EvaluationService services.IEvaluationService
}

func NewUserController(userService services.IUsersService, evaluationService services.IEvaluationService) *UsersController {
	return &UsersController{ UsersService: userService, EvaluationService: evaluationService }
}

func (uc *UsersController) GetUserByID(res http.ResponseWriter, req *http.Request) {
	user, err := uc.UsersService.GetUserByID(mux.Vars(req)["id"])
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{ "error": "Internal server error." })
		return
	}
	if user == nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]string{ "error": "User is not found." })
		return
	}
	disrespectedUser := models.DisrespectedUser{
		User: models.User{ ID: user.ID, Email: user.Email },
		Evaluation: uc.EvaluationService.GenEvaluation(),
	}
	json.NewEncoder(res).Encode(disrespectedUser)
}

func (uc *UsersController) AddUser(res http.ResponseWriter, req *http.Request) {
	data := &models.User{}
	if json.NewDecoder(req.Body).Decode(data) != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{ "error": "Invalid JSON." })
		return
	}
	validationErrors, err := uc.UsersService.AddUser(data)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{ "error": "Internal server error." })
		return
	}
	if len(validationErrors) > 0 {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string][]string{ "errors": validationErrors })
		return
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(data)
}