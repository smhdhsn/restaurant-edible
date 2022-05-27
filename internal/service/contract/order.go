package contract

// OrderService is the interface that order service must implement.
type OrderService interface {
	OrderFood(foodID uint32) (bool, error)
}
