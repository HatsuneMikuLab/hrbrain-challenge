package main

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	connect()
}

func connect2db(url) (*sql.DB, error) {
	//fmt.Println(sql.Drivers())
 	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}