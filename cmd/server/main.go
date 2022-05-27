package main

import (
	"github.com/smhdhsn/restaurant-edible/internal/config"
	"github.com/smhdhsn/restaurant-edible/internal/db"
	"github.com/smhdhsn/restaurant-edible/internal/model"
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
	if err := db.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	// instantiate models.
	fModel := new(model.Food)
	iModel := new(model.Inventory)
	cModel := new(model.Component)

	// instantiate repositories.
	fRepo := mysql.NewFoodRepository(dbConn, *fModel)
	iRepo := mysql.NewInventoryRepository(dbConn, *iModel)
	cRepo := mysql.NewComponentRepository(dbConn, *cModel)

	// instantiate services.
	mServ := service.NewMenuService(fRepo)
	rServ := service.NewRecipeService(fRepo)
	iServ := service.NewInventoryService(iRepo, cRepo, fRepo)

	// instantiate handlers.
	mHandler := handler.NewMenuHandler(mServ)
	rHandler := handler.NewRecipeHandler(rServ)
	iHandler := handler.NewInventoryHandler(iServ)

	// instantiate resources.
	eRes := resource.NewEdibleResource(mHandler, rHandler, iHandler)

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
