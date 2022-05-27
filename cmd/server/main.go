package main

import (
	"github.com/smhdhsn/restaurant-menu/internal/config"
	"github.com/smhdhsn/restaurant-menu/internal/db"
	"github.com/smhdhsn/restaurant-menu/internal/model"
	"github.com/smhdhsn/restaurant-menu/internal/repository/mysql"
	"github.com/smhdhsn/restaurant-menu/internal/server"
	"github.com/smhdhsn/restaurant-menu/internal/server/handler"
	"github.com/smhdhsn/restaurant-menu/internal/server/resource"
	"github.com/smhdhsn/restaurant-menu/internal/service"

	log "github.com/smhdhsn/restaurant-menu/internal/logger"
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
	fRepo := mysql.NewFoodRepository(dbConn, *fModel)
	_ = mysql.NewInventoryRepository(dbConn, *iModel)

	// instantiate services.
	m := service.NewMenuService(fRepo)

	/*
		recipe
			CreateRecipe(fList []*model.Food) error

		menu
			GetFoods() ([]*model.Food, error) -> GetMenu()

		inventory
			BuyComponents(*BuyComponentsReq) error -> Buy()
			Recycle(iRepoContract.RecycleReq) error
			OrderFood(foodID uint32) (bool, error) -> Use()
	*/

	// instantiate handlers.
	mSourceHandler := handler.NewMenuSourceHandler(m)

	// instantiate resources.
	mRes := resource.NewMenuResource(mSourceHandler)

	// instantiate gRPC server.
	s, err := server.New(&conf.Server, mRes)
	if err != nil {
		log.Fatal(err)
	}

	// listen and serve.
	if err := s.Listen(&conf.Server); err != nil {
		log.Fatal(err)
	}
}
