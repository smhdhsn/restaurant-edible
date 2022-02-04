package mysql

import (
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

func (r *FoodRepo) Find() {

}
