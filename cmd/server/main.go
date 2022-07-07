package main

import (
	"github.com/smhdhsn/restaurant-edible/internal/config"
	"github.com/smhdhsn/restaurant-edible/internal/db"
	"github.com/smhdhsn/restaurant-edible/internal/repository/mysql"
	"github.com/smhdhsn/restaurant-edible/internal/server"
	"github.com/smhdhsn/restaurant-edible/internal/server/handler"
	"github.com/smhdhsn/restaurant-edible/internal/server/resource"
	"github.com/smhdhsn/restaurant-edible/internal/service"

	log "github.com/smhdhsn/restaurant-edible/internal/logger"
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
	if err := mysql.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	// instantiate repositories.
	cRepo := mysql.NewComponentRepository(dbConn)
	iRepo := mysql.NewInventoryRepository(dbConn)
	fRepo := mysql.NewFoodRepository(dbConn)

	// instantiate services.
	iServ := service.NewInventoryService(iRepo, cRepo, fRepo)
	rServ := service.NewRecipeService(fRepo)
	mServ := service.NewMenuService(fRepo)

	// instantiate handlers.
	iHandler := handler.NewInventoryHandler(iServ)
	rHandler := handler.NewRecipeHandler(rServ)
	mHandler := handler.NewMenuHandler(mServ)

	// instantiate resources.
	eRes := resource.NewEdibleResource(iHandler, rHandler, mHandler)

	// instantiate gRPC server.
	s, err := server.New(&conf.Server, eRes)
	if err != nil {
		log.Fatal(err)
	}

	// listen and serve.
	if err := s.Listen(); err != nil {
		log.Fatal(err)
	}
}
