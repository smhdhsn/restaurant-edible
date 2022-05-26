package inventory

import (
	"time"

	iRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/inventory"
)

// InventoryService is the interface that inventory service must implement.
type InventoryService interface {
	BuyComponents(*BuyComponentsReq) error
	Recycle(iRepoContract.RecycleReq) error
}

// BuyComponentsReq holds the details of buyed stocks.
type BuyComponentsReq struct {
	StockAmount uint32
	BestBefore  time.Time
	ExpiresAt   time.Time
}
