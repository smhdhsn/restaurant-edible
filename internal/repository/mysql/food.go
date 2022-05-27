package mysql

import (
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-edible/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
)

// FoodRepo contains repository's database connection.
type FoodRepo struct {
	model model.Food
	db    *gorm.DB
}

// NewFoodRepo creates an instance of the repository with database connection.
func NewFoodRepository(db *gorm.DB, m model.Food) repositoryContract.FoodRepository {
	return &FoodRepo{
		model: m,
		db:    db,
	}
}

// GetAvailable gets foods that their components are available.
func (r *FoodRepo) GetAvailable() ([]*model.FoodDTO, error) {
	result := make([]*model.FoodDTO, 0)

	tx := r.db.
		Model(r.model).
		Where(
			"f.id NOT IN (?)",
			unavailableFoods(r.db),
		).
		Find(&result)

	return result, tx.Error
}

// BatchInsert is responsible for storing a chunk of data inside database.
func (r *FoodRepo) BatchInsert(fListDTO []*model.FoodDTO) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, f := range fListDTO {
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
