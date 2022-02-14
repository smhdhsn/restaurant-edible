package mysql

import (
	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository"
	"gorm.io/gorm"
)

// the amount of items being used with every order submittion.
const decrBy = 1

// InventoryRepo contains repository's database connection.
type InventoryRepo struct {
	db *gorm.DB
}

// NewInventoryRepo creates an instance of the repository with database connection.
func NewInventoryRepo(db *gorm.DB) repository.InventoryRepository {
	return &InventoryRepo{db}
}

// UseStocks decreases food components' stock from inventory.
func (s *InventoryRepo) UseStocks(foodID uint) error {
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

// BuyStocks is responsible for buying food components for the inventory, if components' stock are finished or expired.
func (s *InventoryRepo) BuyStocks(iList []*model.Inventory) error {
	return s.db.Model(&model.Inventory{}).CreateInBatches(iList, 100).Error
}
