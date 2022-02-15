package mysql

import (
	"time"

	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository"
	"gorm.io/gorm"
)

// FoodRepo contains repository's database connection.
type FoodRepo struct {
	db *gorm.DB
}

// NewFoodRepo creates an instance of the repository with database connection.
func NewFoodRepo(db *gorm.DB) repository.FoodRepository {
	return &FoodRepo{db}
}

// GetAvailableMeals gets foods that their components are available (not expired or finished).
func (r *FoodRepo) GetAvailableMeals() ([]*model.Food, error) {
	result := make([]*model.Food, 0)

	tx := r.db.
		Table("foods").
		Where(
			"foods.id NOT IN (?)",
			r.db.
				Table("food_components").
				Select("food_components.food_id").
				Where(
					"food_components.component_id NOT IN (?)",
					r.db.Table("inventories").Select("inventories.component_id"),
				).
				Or(
					"food_components.component_id IN (?)",
					r.db.
						Table("inventories").
						Select("inventories.component_id").
						Where("inventories.expires_at < ?", time.Now()).
						Or("inventories.stock = ?", 0),
				),
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
