package main

import (
	"fmt"
	"log"

	"github.com/smhdhsn/food/internal/config"
	"github.com/smhdhsn/food/internal/db"
)

// main is the main application entry.
func main() {
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.Connect(conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf)
}
