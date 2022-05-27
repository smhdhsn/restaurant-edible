package contract

import (
	"time"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
)

// InventoryService is the interface that inventory service must implement.
type InventoryService interface {
	BuyComponents(*BuyComponentsReq) error
	Recycle(repositoryContract.RecycleReq) error
}

// BuyComponentsReq holds the details of buyed stocks.
type BuyComponentsReq struct {
	StockAmount uint32
	BestBefore  time.Time
	ExpiresAt   time.Time
}
