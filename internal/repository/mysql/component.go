package mysql

import (
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
