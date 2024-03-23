package order

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftShopOrder struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	ShopOrderData   `bson:",inline" json:",inline"`
}

type ShopOrderData struct {
	Name      string  `bson:"name" json:"name,omitempty"`
	SortOrder float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note      string  `bson:"note" json:"note"`
}

func NewShopOrder(data ShopOrderData) *GftShopOrder {
	item := &GftShopOrder{
		ShopOrderData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
