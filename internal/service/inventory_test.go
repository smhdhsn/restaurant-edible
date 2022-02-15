package service

import (
	"testing"
	"time"

	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository/mock"
	"github.com/stretchr/testify/assert"
)

// Random data for testing.
var (
	randAmount     = uint(5)
	randBestBefore = time.Now().AddDate(0, 3, 0)
	randExpiresAt  = time.Now().AddDate(0, 5, 0)
)

func TestBuyComponents(t *testing.T) {
	// Arrange
	c := model.Component{
		Title: "randomTitle",
	}

	cRepoMock := new(mock.ComponentRepo)
	cRepoMock.On("GetUnavailable").Return([]*model.Component{&c}, nil)

	i := model.Inventory{
		ComponentID: c.ID,
		Stock:       randAmount,
		BestBefore:  randBestBefore,
		ExpiresAt:   randExpiresAt,
	}
	iRepoMock := new(mock.InventoryRepo)
	iRepoMock.On("BuyStocks", []*model.Inventory{&i}).Return(nil)

	sut := NewInventoryService(iRepoMock, cRepoMock)

	// Act
	err := sut.BuyComponents(&BuyComponentsReq{
		StockAmount: randAmount,
		BestBefore:  randBestBefore,
		ExpiresAt:   randExpiresAt,
	})

	// Assert
	assert.NoError(t, err)
}
