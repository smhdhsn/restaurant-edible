package main

import (
	"github.com/smhdhsn/restaurant-menu/internal/config"
	"github.com/smhdhsn/restaurant-menu/internal/db"
	"github.com/smhdhsn/restaurant-menu/internal/http"
	"github.com/smhdhsn/restaurant-menu/internal/model"
	"github.com/smhdhsn/restaurant-menu/internal/repository/mysql"

	log "github.com/smhdhsn/restaurant-menu/internal/logger"
	mServ "github.com/smhdhsn/restaurant-menu/internal/service/menu"
	oServ "github.com/smhdhsn/restaurant-menu/internal/service/order"
)

// main is the main application entry.
func main() {
	// read configurations.
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	// create a database connection.
	dbConn, err := db.Connect(&conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	// initialize auto migration.
	if err := db.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	// instantiate models.
	fModel := new(model.Food)
	iModel := new(model.Inventory)

	// instantiate repositories.
	fRepo := mysql.NewFoodRepo(dbConn, *fModel)
	iRepo := mysql.NewInventoryRepo(dbConn, *iModel)

	// instantiate services.
	m := mServ.NewMenuServ(fRepo)
	o := oServ.NewOrderServ(fRepo, iRepo)

	// instantiate handlers.

	// instantiate resources.

	// instantiate gRPC server.
	s, err := http.New(&conf.Server, m, o)
	if err != nil {
		log.Fatal(err)
	}

	// listen and serve.
	if err := s.Listen(&conf.Server); err != nil {
		log.Fatal(err)
	}
}
