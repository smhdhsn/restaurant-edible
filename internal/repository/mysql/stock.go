package mysql

import (
	"github.com/smhdhsn/food/internal/repository"
	"gorm.io/gorm"
)

// decrBy holds the amount of items being used with every order submittion.
const decrBy = 1

// StockRepo contains repository's database connection.
type StockRepo struct {
	db *gorm.DB
}

// NewStockRepo creates an instance of the repository with database connection.
func NewStockRepo(db *gorm.DB) repository.StockRepository {
	return &StockRepo{db}
}

// UseIngredients decreases the stock amount of ingredients related to a food.
func (s *StockRepo) UseIngredients(foodID uint) error {
	tx := s.db

	tx.
		Table("stocks").
		Where(
			"stocks.ingredient_id IN (?)",
			tx.
				Table("food_ingredients").
				Select("food_ingredients.ingredient_id").
				Where("food_ingredients.food_id = ?", foodID),
		).
		Update("stocks.stock", gorm.Expr("stocks.stock - ?", decrBy))

	return tx.Error
}
