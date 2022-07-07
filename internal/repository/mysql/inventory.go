package mysql

import (
	"time"

	"gorm.io/gorm"

	"github.com/pkg/errors"
	"github.com/smhdhsn/restaurant-edible/internal/repository/entity"

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

// Recycle is responsible for cleaning up inventory from useless items.
func (r *InventoryRepo) Recycle(rEntity *entity.Recycle) error {
	err := r.db.
		Table("inventories AS i").
		Where("i.expires_at < ? AND ?", time.Now(), rEntity.Expired).
		Or("i.stock = 0  AND ?", rEntity.Finished).
		Delete(r.model).
		Error

	if err != nil {
		return errors.Wrap(err, "error on recycling inventory items inside database")
	}

	return nil
}

// batchSize holds the size of every batch to be sent to database to be saved.
const batchSize = 100

// Buy is responsible for buying food components for the inventory, if components' stock are finished or expired.
func (r *InventoryRepo) Buy(iListEntity []*entity.Inventory) error {
	err := r.db.
		Model(r.model).
		CreateInBatches(iListEntity, batchSize).
		Error

	if err != nil {
		return errors.Wrap(err, "error on buying food components' inventory stock inside database")
	}

	return nil
}

// the amount of items being used with every order submittion.
const decrBy = 1

// Use decreases food components' stock from inventory.
func (r *InventoryRepo) Use(fEntity *entity.Food) error {
	err := r.db.
		Table("inventories AS i").
		Where("i.expires_at > ?", time.Now()).
		Where("i.stock > 0").
		Where(
			"i.component_id IN (?)",
			componentsOfFood(r.db, fEntity.ID),
		).
		Update("i.stock", gorm.Expr("i.stock - ?", decrBy)).
		Error

	if err != nil {
		return errors.Wrap(err, "error on decreasing food components' inventory stock inside database")
	}

	return nil
}

// availableInventoryItems is the subquery responsible for getting available food components from inventory.
func availableInventoryItems(db *gorm.DB) *gorm.DB {
	return db.
		Table("inventories AS i").
		Select("i.component_id").
		Where("i.expires_at > ?", time.Now()).
		Where("i.stock > 0")
}
