package service

import (
	"testing"

	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository"
	"github.com/smhdhsn/food/internal/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestBuyComponents(t *testing.T) {
	// Arrange
	c := model.Component{
		Title: randStr,
	}

	cRepoMock := new(mock.ComponentRepo)
	cRepoMock.On("GetUnavailable").Return([]*model.Component{&c}, nil)

	i := model.Inventory{
		ComponentID: c.ID,
		Stock:       randUINT,
		BestBefore:  randDate,
		ExpiresAt:   randDate,
	}
	iRepoMock := new(mock.InventoryRepo)
	iRepoMock.On("Buy", []*model.Inventory{&i}).Return(nil)

	sut := NewInventoryService(iRepoMock, cRepoMock)

	// Act
	err := sut.BuyComponents(&BuyComponentsReq{
		StockAmount: randUINT,
		BestBefore:  randDate,
		ExpiresAt:   randDate,
	})

	// Assert
	assert.NoError(t, err)
}

func TestRecycle(t *testing.T) {
	// Arrange
	req := repository.RecycleReq{
		Finished: randBool,
		Expired:  randBool,
	}

	iRepoMock := new(mock.InventoryRepo)
	iRepoMock.On("Clean", req).Return(nil)

	sut := NewInventoryService(iRepoMock, nil)

	// Act
	err := sut.Recycle(req)

	// Assert
	assert.NoError(t, err)
}
