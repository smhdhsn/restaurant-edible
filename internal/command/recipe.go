package command

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/smhdhsn/food/internal/config"
	"github.com/smhdhsn/food/internal/db"
	"github.com/smhdhsn/food/internal/repository/mysql"
	"github.com/smhdhsn/food/internal/service"
	"github.com/spf13/cobra"
)

// recipeCMD is the subcommands responsible for storing sample data inside database.
var recipeCMD = &cobra.Command{
	Use:   "recipe",
	Short: "Stores sample recipes inside database.",
	Run:   seedRun,
}

// init function will be executed when this package is called.
func init() {
	rootCMD.AddCommand(recipeCMD)

	recipeCMD.Flags().StringP("json", "j", "", "Path to recipe JSON.")
	recipeCMD.MarkFlagRequired("json")
}

// seedRun is responsible for running the script.
func seedRun(cmd *cobra.Command, args []string) {
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
	cRepo := mysql.NewComponentRepo(dbConn)

	rService := service.NewRecipeService(fRepo, cRepo)

	j, err := cmd.Flags().GetString("json")
	if err != nil {
		log.Fatal(err)
	}

	data, err := getFromJsonFile(j)
	if err != nil {
		log.Fatal(err)
	}

	err = rService.CreateRecipe(data)
	if err != nil {
		log.Fatal(err)
	}
}

// getFromJsonFile is responsible for getting recipe schema from JSON file.
func getFromJsonFile(path string) (*service.RecipeSchema, error) {
	j, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	defer j.Close()

	b, err := ioutil.ReadAll(j)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read bytes")
	}

	var data service.RecipeSchema
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal json")
	}

	return &data, nil
}
