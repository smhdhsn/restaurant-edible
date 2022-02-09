package mysql

import (
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
