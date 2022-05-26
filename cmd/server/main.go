package main

import (
	"github.com/smhdhsn/restaurant-menu/internal/config"
	"github.com/smhdhsn/restaurant-menu/internal/db"
	"github.com/smhdhsn/restaurant-menu/internal/http"
	"github.com/smhdhsn/restaurant-menu/internal/repository/mysql"
	"github.com/smhdhsn/restaurant-menu/internal/service"

	log "github.com/smhdhsn/restaurant-menu/internal/logger"
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

	fRepo := mysql.NewFoodRepo(dbConn)
	iRepo := mysql.NewInventoryRepo(dbConn)

	mService := service.NewMenuService(fRepo)
	oService := service.NewOrderService(fRepo, iRepo)

	httpServer, err := http.New(mService, oService)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(httpServer.Listen(conf.Server.Host, conf.Server.Port))
}
