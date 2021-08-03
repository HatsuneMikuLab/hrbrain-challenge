package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func Connect2db(url string) *sql.DB {
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