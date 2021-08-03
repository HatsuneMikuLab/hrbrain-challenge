package main

import (
	"github.com/HatsuneMikuLab/hrbrain-challenge/db"
	"time"
	"os"
	"net/http"
	"database/sql"
	"log"
	"github.com/gorilla/mux"
	"github.com/HatsuneMikuLab/hrbrain-challenge/services"
	"github.com/HatsuneMikuLab/hrbrain-challenge/controllers"
	"github.com/HatsuneMikuLab/hrbrain-challenge/middlewares"
	
)

func main() {
	log.Println(os.Args[2])
	db := db.Connect2db(os.Args[2])
	
	router := InitRouter(db)
	http.ListenAndServe(os.Args[1], router)
}


func InitRouter(db *sql.DB) http.Handler {
	router := mux.NewRouter()
	userCtrl := controllers.NewUserController(
		services.NewUserService(db), 
		services.NewEvaluationService(),
		services.NewCacheService(),
	)
	router.Use(middlewares.SetHeadersMiddleware)
	router.HandleFunc("/api/users", userCtrl.AddUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", userCtrl.GetUserByID).Methods("GET")
	
	return http.TimeoutHandler(router, 2 * time.Second, "Request timeout")
}