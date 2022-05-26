package mysql

import (
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-menu/internal/model"

	fRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/food"
)

// FoodRepo contains repository's database connection.
type FoodRepo struct {
	model model.Food
	db    *gorm.DB
}

// NewFoodRepo creates an instance of the repository with database connection.
func NewFoodRepo(db *gorm.DB, m model.Food) fRepoContract.FoodRepository {
	return &FoodRepo{
		model: m,
		db:    db,
	}
}

// GetAvailable gets foods that their components are available.
func (r *FoodRepo) GetAvailable() ([]*model.Food, error) {
	result := make([]*model.Food, 0)

	tx := r.db.
		Table("foods AS f").
		Where(
			"f.id NOT IN (?)",
			unavailableFoods(r.db),
		).
		Find(&result)

	return result, tx.Error
}

// BatchInsert is responsible for storing a chunk of data inside database.
func (r *FoodRepo) BatchInsert(fList []*model.Food) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, f := range fList {
			for _, c := range f.Components {
				if err := tx.Where(model.Component{Title: c.Title}).FirstOrCreate(&c).Error; err != nil {
					return err
				}
			}

			if err := tx.Where(model.Food{Title: f.Title}).FirstOrCreate(&f).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
