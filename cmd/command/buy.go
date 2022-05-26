package main

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/smhdhsn/restaurant-menu/internal/config"
	"github.com/smhdhsn/restaurant-menu/internal/db"
	"github.com/smhdhsn/restaurant-menu/internal/model"
	"github.com/smhdhsn/restaurant-menu/internal/repository/mysql"

	log "github.com/smhdhsn/restaurant-menu/internal/logger"
	iServContract "github.com/smhdhsn/restaurant-menu/internal/service/contract/inventory"
	iServ "github.com/smhdhsn/restaurant-menu/internal/service/inventory"
)

var (
	amount     = uint32(3)                   // the amount of items being added to inventory with every buy.
	bestBefore = time.Now().AddDate(0, 2, 0) // item's best usage time.
	expiresAt  = time.Now().AddDate(0, 5, 0) // item's expiration time.
)

// buyCMD is the subcommands responsible for creating food components.
var buyCMD = &cobra.Command{
	Use:   "buy",
	Short: "Stores new food components inside database if their components' stock are finished or expired.",
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
		cModel := new(model.Component)

		// instantiate repositories.
		iRepo := mysql.NewInventoryRepo(dbConn, *iModel)
		cRepo := mysql.NewComponentRepo(dbConn, *cModel)

		// instantiate services.
		i := iServ.NewInventoryServ(iRepo, cRepo)

		// read amount from cli.
		a, err := cmd.Flags().GetUint32("amount")
		if err != nil {
			log.Fatal(err)
		}

		// create service request.
		req := iServContract.BuyComponentsReq{
			StockAmount: a,
			BestBefore:  bestBefore,
			ExpiresAt:   expiresAt,
		}

		// call service.
		err = i.BuyComponents(&req)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// init function will be executed when this package is called.
func init() {
	rootCMD.AddCommand(buyCMD)

	buyCMD.Flags().Uint32P("amount", "a", amount, "The amount of stocks added to inventory after each buy.")
}
