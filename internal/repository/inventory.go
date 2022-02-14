package repository

import "github.com/smhdhsn/food/internal/model"

// InventoryRepository is the interface representing inventory repository or it's mock.
type InventoryRepository interface {
	UseStocks(uint) error
	BuyStocks([]*model.Inventory) error
}
