package repository

// InventoryRepository is the interface representing inventory repository or it's mock.
type InventoryRepository interface {
	UseComponents(uint) error
}
