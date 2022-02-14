package mysql

import (
	"github.com/smhdhsn/food/internal/repository"
	"gorm.io/gorm"
)

// decrBy holds the amount of items being used with every order submittion.
const decrBy = 1

// InventoryRepo contains repository's database connection.
type InventoryRepo struct {
	db *gorm.DB
}

// NewInventoryRepo creates an instance of the repository with database connection.
func NewInventoryRepo(db *gorm.DB) repository.InventoryRepository {
	return &InventoryRepo{db}
}

// UseComponentsFor decreases food components' stock from inventory.
func (s *InventoryRepo) UseComponents(foodID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return tx.
			Table("inventories").
			Where(
				"inventories.component_id IN (?)",
				tx.
					Table("food_components").
					Select("food_components.component_id").
					Where("food_components.food_id = ?", foodID),
			).
			Update("inventories.stock", gorm.Expr("inventories.stock - ?", decrBy)).Error
	})
}
