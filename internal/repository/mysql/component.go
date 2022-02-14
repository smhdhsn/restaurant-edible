package mysql

import (
	"time"

	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository"
	"gorm.io/gorm"
)

// ComponentRepo contains repository's database connection.
type ComponentRepo struct {
	db *gorm.DB
}

// NewComponentRepo creates an instance of the repository with database connection.
func NewComponentRepo(db *gorm.DB) repository.ComponentRepository {
	return &ComponentRepo{db}
}

// GetUnavailable is responsible for getting food components that are not avaiable.
// Components with finished or expired stocks are counted as unavailable.
func (r *ComponentRepo) GetUnavailable() ([]*model.Component, error) {
	result := make([]*model.Component, 0)

	tx := r.db.
		Table("components").
		Where(
			"components.id IN (?)",
			r.db.Table("inventories").Select("inventories.component_id").Where("inventories.stock = ?", 0).Or("inventories.expires_at < ?", time.Now()),
		).
		Or(
			"components.id NOT IN (?)",
			r.db.Table("inventories").Select("inventories.component_id"),
		).
		Find(&result)

	return result, tx.Error
}
