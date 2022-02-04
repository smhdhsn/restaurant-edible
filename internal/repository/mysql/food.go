package mysql

import (
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

	tx := r.db.
		Joins("JOIN food_ingredients ON foods.id = food_ingredients.food_id").
		Joins("JOIN ingredients ON food_ingredients.ingredient_id = ingredients.id").
		Joins("JOIN stocks ON ingredients.id = stocks.ingredient_id").
		Group("foods.id").
		Find(&result)

	return result, tx.Error
}
