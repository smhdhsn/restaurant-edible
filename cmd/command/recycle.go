package main

import (
	"github.com/spf13/cobra"

	"github.com/smhdhsn/restaurant-menu/internal/config"
	"github.com/smhdhsn/restaurant-menu/internal/db"
	"github.com/smhdhsn/restaurant-menu/internal/model"
	"github.com/smhdhsn/restaurant-menu/internal/repository/mysql"

	log "github.com/smhdhsn/restaurant-menu/internal/logger"
	iRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/inventory"
	iServ "github.com/smhdhsn/restaurant-menu/internal/service/inventory"
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
		iRepo := mysql.NewInventoryRepo(dbConn, *iModel)

		// instantiate services.
		i := iServ.NewInventoryService(iRepo, nil)

		// create service request.
		req := iRepoContract.RecycleReq{
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
