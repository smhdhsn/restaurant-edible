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

// GetAvailableMeals gets foods that their ingredients are available (not expired|finished).
func (r *FoodRepo) GetAvailableMeals() ([]*model.Food, error) {
	result := make([]*model.Food, 0)

	tx := r.db

	tx.
		Table("foods").
		Where(
			"foods.id NOT IN (?)",
			tx.
				Table("food_ingredients").
				Select("food_ingredients.food_id").
				Where(
					"food_ingredients.ingredient_id NOT IN (?)",
					tx.Table("stocks").Select("stocks.ingredient_id"),
				).
				Or(
					"food_ingredients.ingredient_id IN (?)",
					tx.
						Table("stocks").
						Select("stocks.ingredient_id").
						Where("stocks.expires_at < ?", time.Now()).
						Or("stocks.stock = ?", 0),
				),
		).
		Find(&result)

	return result, tx.Error
}
