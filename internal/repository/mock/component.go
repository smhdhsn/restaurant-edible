package mock

import (
	"github.com/smhdhsn/food/internal/model"
	"github.com/stretchr/testify/mock"
)

// ComponentRepo is component repository's mock.
type ComponentRepo struct {
	mock.Mock
}

// GetUnavailable is a mocked method in mocked component repository.
func (r *ComponentRepo) GetUnavailable() ([]*model.Component, error) {
	args := r.Mock.Called()

	if v := args.Get(0); v != nil {
		return v.([]*model.Component), nil
	}

	return nil, args.Error(1)
}
