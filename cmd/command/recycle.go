package main

import (
	"log"

	"github.com/smhdhsn/food/internal/config"
	"github.com/smhdhsn/food/internal/db"
	"github.com/smhdhsn/food/internal/repository"
	"github.com/smhdhsn/food/internal/repository/mysql"
	"github.com/smhdhsn/food/internal/service"
	"github.com/spf13/cobra"
)

// This section holds the items to be cleaned up from inventory.
var (
	recycleFinished = false
	recycleExpired  = false
)

// recycleCMD is the subcommands responsible for cleaning up inventory from unusable items.
var recycleCMD = &cobra.Command{
	Use:   "recycle",
	Short: "Deletes useless items from inventory.",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.LoadConf()
		if err != nil {
			log.Fatal(err)
		}

		dbConn, err := db.Connect(conf.DB)
		if err != nil {
			log.Fatal(err)
		}

		iRepo := mysql.NewInventoryRepo(dbConn)

		iService := service.NewInventoryService(iRepo, nil)

		err = iService.Recycle(repository.RecycleReq{
			Finished: recycleFinished,
			Expired:  recycleExpired,
		})
		if err != nil {
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
