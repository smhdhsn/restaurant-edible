package mysql

import (
	"time"

	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-menu/internal/model"
	"github.com/smhdhsn/restaurant-menu/internal/repository"
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

// Buy is responsible for buying food components for the inventory, if components' stock are finished or expired.
func (r *InventoryRepo) Buy(iList []*model.Inventory) error {
	return r.db.Model(&model.Inventory{}).CreateInBatches(iList, 100).Error
}

// Clean is responsible for cleaning up inventory from useless items.
func (r *InventoryRepo) Clean(req repository.RecycleReq) error {
	return r.db.
		Table("inventories AS i").
		Where("i.expires_at < ? AND ?", time.Now(), req.Expired).
		Or("i.stock = 0  AND ?", req.Finished).
		Delete(&model.Inventory{}).
		Error
}

// Use decreases food components' stock from inventory.
func (r *InventoryRepo) Use(foodID uint) error {
	return r.db.
		Table("inventories AS i").
		Where("i.expires_at > ?", time.Now()).
		Where("i.stock > 0").
		Where(
			"i.component_id IN (?)",
			componentsOfFood(r.db, foodID),
		).
		Update("i.stock", gorm.Expr("i.stock - ?", decrBy)).
		Error
}

// availableInventoryItems is the subquery responsible for getting available food components from inventory.
func availableInventoryItems(db *gorm.DB) *gorm.DB {
	return db.
		Table("inventories AS i").
		Select("i.component_id").
		Where("i.expires_at > ?", time.Now()).
		Where("i.stock > 0")
}
