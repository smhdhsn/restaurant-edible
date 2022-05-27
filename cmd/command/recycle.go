package main

import (
	"github.com/spf13/cobra"

	"github.com/smhdhsn/restaurant-edible/internal/config"
	"github.com/smhdhsn/restaurant-edible/internal/db"
	"github.com/smhdhsn/restaurant-edible/internal/model"
	"github.com/smhdhsn/restaurant-edible/internal/repository/mysql"
	"github.com/smhdhsn/restaurant-edible/internal/service"

	log "github.com/smhdhsn/restaurant-edible/internal/logger"
	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
)

// This section holds the items to be cleaned up from inventory.
var recycleFinished, recycleExpired bool

// recycleCMD is the subcommands responsible for cleaning up inventory from unusable items.
var recycleCMD = &cobra.Command{
	Use:   "recycle",
	Short: "Deletes useless items from inventory.",
	Run: func(cmd *cobra.Command, args []string) {
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
		iModel := new(model.Inventory)

		// instantiate repositories.
		iRepo := mysql.NewInventoryRepository(dbConn, *iModel)

		// instantiate services.
		i := service.NewInventoryService(iRepo, nil)

		// create service request.
		req := repositoryContract.RecycleReq{
			Finished: recycleFinished,
			Expired:  recycleExpired,
		}

		// call service.
		if err := i.Recycle(req); err != nil {
			log.Fatal(err)
		}
	},
}

// init function will be executed when this package is called.
func init() {
	rootCMD.AddCommand(recycleCMD)

	recycleCMD.Flags().BoolVarP(&recycleFinished, "finished", "f", false, "Recycle finished items.")
	recycleCMD.Flags().BoolVarP(&recycleExpired, "expired", "e", false, "Recycle expired items.")
}
