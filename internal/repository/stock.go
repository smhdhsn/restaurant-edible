package repository

import "errors"

// StockRepository is the interface representing stock repository or it's mock.
type StockRepository interface {
	UseIngredients([]uint) error
}

// ErrNotAvailable occurs when an order can not be fulfilled because of the lack of ingredients.
var ErrNotAvailable = errors.New("requested order cannot be fulfilled because of the lack of ingredients")
