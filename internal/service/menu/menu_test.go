package menu

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smhdhsn/restaurant-menu/internal/model"
	"github.com/smhdhsn/restaurant-menu/internal/repository/mock"
	"github.com/smhdhsn/restaurant-menu/util/random"
)

func TestGetFoods(t *testing.T) {
	// Arrange
	randTitle := random.GenerateString(5)

	f := model.Food{
		Title: randTitle,
	}
	fRepoMock := new(mock.FoodRepo)
	fRepoMock.On("GetAvailable").Return([]*model.Food{&f}, nil)

	sut := NewMenuService(fRepoMock)

	// Act
	fList, err := sut.GetFoods()

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, fList, &f)
}
