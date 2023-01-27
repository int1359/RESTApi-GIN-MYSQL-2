package main

import (
	"database/sql"
	"log"
	"rest/api/config"
	"rest/api/routes"
)

var err error

func init() {
	config.DB, err = sql.Open("mysql", config.DbURL())

	if err != nil {
		log.Fatal("Status:", err)
	}

}

func main() {

	r := routes.SetupRouter()
	r.Run()
}
