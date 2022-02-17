package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/smhdhsn/food/internal/config"
	"github.com/smhdhsn/food/internal/db"
	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository/mysql"
	"github.com/smhdhsn/food/internal/service"
	"github.com/spf13/cobra"
)

// RecipeSchema holds schema for recipe JSON.
type RecipeFileSchema struct {
	Foods []struct {
		Title      string   `json:"title"`
		Components []string `json:"components"`
	} `json:"foods"`
}

// recipeCMD is the subcommands responsible for storing sample data inside database.
var recipeCMD = &cobra.Command{
	Use:   "recipe",
	Short: "Stores sample recipes inside database.",
	Run: func(cmd *cobra.Command, args []string) {
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

		rService := service.NewRecipeService(fRepo)

		j, err := cmd.Flags().GetString("json")
		if err != nil {
			log.Fatal(err)
		}

		data, err := getFromFile(j)
		if err != nil {
			log.Fatal(err)
		}

		err = rService.CreateRecipe(data)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// init function will be executed when this package is called.
func init() {
	rootCMD.AddCommand(recipeCMD)

	recipeCMD.Flags().StringP("json", "j", "", "Path to recipe JSON.")
	recipeCMD.MarkFlagRequired("json")
}

// getFromFile is responsible for getting recipe schema from JSON file.
func getFromFile(path string) ([]*model.Food, error) {
	j, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	defer j.Close()

	b, err := ioutil.ReadAll(j)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read bytes")
	}

	var schema RecipeFileSchema
	err = json.Unmarshal(b, &schema)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal json")
	}

	m := schema.convertToModel()

	return m, nil
}

// convertToModel is responsible for converting schema to application's models.
func (r *RecipeFileSchema) convertToModel() []*model.Food {
	cList := make([]*model.Component, 0)
	fList := make([]*model.Food, 0)
	for _, f := range r.Foods {
		for _, cTitle := range f.Components {
			cList = append(cList, &model.Component{Title: cTitle})
		}

		fList = append(fList, &model.Food{Title: f.Title, Components: cList})
	}

	return fList
}
