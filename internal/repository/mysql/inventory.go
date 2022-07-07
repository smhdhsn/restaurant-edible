package mysql

import (
	"time"

	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-edible/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
)

// inventory represents the inventories table on database.
type inventory struct {
	ID          uint32    `gorm:"primaryKey"`
	ComponentID uint32    `gorm:"index;not null"`
	Component   component `gorm:"constraint:OnDelete:CASCADE"`
	Stock       uint32    `gorm:"not null"`
	ExpiresAt   time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// InventoryRepo contains repository's database connection.
type InventoryRepo struct {
	model inventory
	db    *gorm.DB
}

// NewInventoryRepo creates an instance of the repository with database connection.
func NewInventoryRepository(db *gorm.DB) repositoryContract.InventoryRepository {
	return &InventoryRepo{
		model: inventory{},
		db:    db,
	}
}

// Buy is responsible for buying food components for the inventory, if components' stock are finished or expired.
func (r *InventoryRepo) Buy(iListDTO []*model.InventoryDTO) error {
	return r.db.Model(r.model).CreateInBatches(iListDTO, 100).Error
}

// Recycle is responsible for cleaning up inventory from useless items.
func (r *InventoryRepo) Recycle(finished, expired bool) error {
	return r.db.
		Table("inventories AS i").
		Where("i.expires_at < ? AND ?", time.Now(), expired).
		Or("i.stock = 0  AND ?", finished).
		Delete(r.model).
		Error
}

// the amount of items being used with every order submittion.
const decrBy = 1

// Use decreases food components' stock from inventory.
func (r *InventoryRepo) Use(fDTO *model.FoodDTO) error {
	return r.db.
		Table("inventories AS i").
		Where("i.expires_at > ?", time.Now()).
		Where("i.stock > 0").
		Where(
			"i.component_id IN (?)",
			componentsOfFood(r.db, fDTO.ID),
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
