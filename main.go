package main

import (
	"database/sql"
	"fmt"
	"log"
	"rest/api/config"
	"rest/api/routes"
)

var err error

func init() {
	config.DB, err = sql.Open("mysql", config.DbURL())
	fmt.Print(config.DB)
	if err != nil {
		log.Fatal("Status:", err)
	}

}

func main() {

	r := routes.SetupRouter()
	r.Run()
}
