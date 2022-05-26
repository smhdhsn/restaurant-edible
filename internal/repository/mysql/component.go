package mysql

import (
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-menu/internal/model"

	cRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/component"
)

// ComponentRepo contains repository's database connection.
type ComponentRepo struct {
	model model.Component
	db    *gorm.DB
}

// NewComponentRepo creates an instance of the repository with database connection.
func NewComponentRepo(db *gorm.DB, m model.Component) cRepoContract.ComponentRepository {
	return &ComponentRepo{
		model: m,
		db:    db,
	}
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
