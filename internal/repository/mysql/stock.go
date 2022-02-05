package mysql

import (
	"github.com/smhdhsn/food/internal/repository"
	"gorm.io/gorm"
)

// StockRepo contains repository's database connection.
type StockRepo struct {
	db *gorm.DB
}

// NewStockRepo creates an instance of the repository with database connection.
func NewStockRepo(db *gorm.DB) repository.StockRepository {
	return &StockRepo{db}
}

// UseIngredients decreases the amount of stock.
func (s *StockRepo) UseIngredients(ingrdIDs []uint) error {
	tx := s.db.Begin()

	for _, ingrdID := range ingrdIDs {
		err := s.use(ingrdID)
		if err != nil {
			tx.Rollback()

			return err
		}
	}

	tx.Commit()

	return nil
}

// use is responsible for decreasing the amount of an ingredient's stock.
func (s *StockRepo) use(ingrdID uint) error {
	tx := s.db.Exec("UPDATE stocks SET stock = stock - 1 WHERE stocks.id = ?", ingrdID)

	return tx.Error
}
