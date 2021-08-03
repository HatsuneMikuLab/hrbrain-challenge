package main

import (
	"time"
	"os"
	"net/http"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/HatsuneMikuLab/hrbrain-challenge/services"
	"github.com/HatsuneMikuLab/hrbrain-challenge/controllers"
	"github.com/HatsuneMikuLab/hrbrain-challenge/middlewares"
	
)

func main() {
	log.Println(os.Args[2])
	db := connect2db(os.Args[2])
	
	router := InitRouter(db)
	http.ListenAndServe(os.Args[1], router)
}

func connect2db(url string) *sql.DB {
	//fmt.Println(sql.Drivers())
 	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
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