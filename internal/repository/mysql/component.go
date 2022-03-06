package mysql

import (
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
func (r *ComponentRepo) GetUnavailable() ([]*model.Component, error) {
	result := make([]*model.Component, 0)

	tx := r.db.
		Table("components AS c").
		Where(
			"c.id NOT IN (?)",
			availableInventoryItems(r.db),
		).
		Find(&result)

	return result, tx.Error
}
