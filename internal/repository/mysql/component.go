package mysql

import (
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-edible/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
)

// ComponentRepo contains repository's database connection.
type ComponentRepo struct {
	model model.Component
	db    *gorm.DB
}

// NewComponentRepo creates an instance of the repository with database connection.
func NewComponentRepository(db *gorm.DB, m model.Component) repositoryContract.ComponentRepository {
	return &ComponentRepo{
		model: m,
		db:    db,
	}
}

// GetUnavailable is responsible for getting food components that are not avaiable.
func (r *ComponentRepo) GetUnavailable() ([]*model.ComponentDTO, error) {
	result := make([]*model.ComponentDTO, 0)

	tx := r.db.
		Model(r.model).
		Table("components AS c").
		Where(
			"c.id NOT IN (?)",
			availableInventoryItems(r.db),
		).
		Find(&result)

	return result, tx.Error
}
