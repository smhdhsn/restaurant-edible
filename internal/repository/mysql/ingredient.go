package mysql

import (
	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository"
	"gorm.io/gorm"
)

// IngredientRepo contains repository's database connection.
type IngredientRepo struct {
	db *gorm.DB
}

// NewIngredientRepo creates an instance of the repository with database connection.
func NewIngredientRepo(db *gorm.DB) repository.IngredientRepository {
	return &IngredientRepo{db}
}

// GetFoodIngredients is responsible for getting a food's ingredients.
func (r *IngredientRepo) GetFoodIngredients(foodID uint) ([]*model.Ingredient, error) {
	result := make([]*model.Ingredient, 0)

	tx := r.db.
		Model(&model.Ingredient{}).
		Joins("JOIN food_ingredients ON food_ingredients.ingredient_id = ingredients.id").
		Where("food_ingredients.food_id = ?", foodID).
		Find(&result)

	return result, tx.Error
}
