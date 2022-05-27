package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smhdhsn/restaurant-menu/internal/model"
	"github.com/smhdhsn/restaurant-menu/internal/repository/mock"
	"github.com/smhdhsn/restaurant-menu/util/random"

	repositoryContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-menu/internal/service/contract"
)

func TestBuyComponents(t *testing.T) {
	// Arrange
	randTitle := random.GenerateString(5)
	randStock := random.GenerateUint32(1, 3)
	randBestBefore := random.GenerateDateBetween(2022, 2025)
	randExpiresAt := random.GenerateDateBetween(2025, 2030)

	c := model.Component{
		Title: randTitle,
	}

	cRepoMock := new(mock.ComponentRepo)
	cRepoMock.On("GetUnavailable").Return([]*model.Component{&c}, nil)

	i := model.Inventory{
		ComponentID: c.ID,
		Stock:       randStock,
		BestBefore:  randBestBefore,
		ExpiresAt:   randExpiresAt,
	}
	iRepoMock := new(mock.InventoryRepo)
	iRepoMock.On("Buy", []*model.Inventory{&i}).Return(nil)

	req := serviceContract.BuyComponentsReq{
		StockAmount: randStock,
		BestBefore:  randBestBefore,
		ExpiresAt:   randExpiresAt,
	}

	sut := NewInventoryService(iRepoMock, cRepoMock)

	// Act
	err := sut.BuyComponents(&req)

	// Assert
	assert.NoError(t, err)
}

func TestRecycle(t *testing.T) {
	// Arrange
	randFinished := random.GenerateBool()
	randExpired := random.GenerateBool()

	req := repositoryContract.RecycleReq{
		Finished: randFinished,
		Expired:  randExpired,
	}

	iRepoMock := new(mock.InventoryRepo)
	iRepoMock.On("Clean", req).Return(nil)

	sut := NewInventoryService(iRepoMock, nil)

	// Act
	err := sut.Recycle(req)

	// Assert
	assert.NoError(t, err)
}
