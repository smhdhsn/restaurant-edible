package db

import (
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-menu/internal/model"
)

// MigrationModels holds the schema of the models to be migrated to database.
var MigrationModels = [...]interface{}{
	&model.Food{},
	&model.Component{},
	&model.Inventory{},
}

// InitMigrations migrates models to the database.
func InitMigrations(db *gorm.DB) error {
	return db.AutoMigrate(MigrationModels[:]...)
}
