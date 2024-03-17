package product

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftShopProduct struct {
	types.Entity    `bson:",inline" json:",inline"`
	ShopProductData `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
}

type ShopProductData struct {
	Name      string  `bson:"name" json:"name,omitempty"`
	SortOrder float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note      string  `bson:"note" json:"note"`
}

func NewShopProduct(data ShopProductData) *GftShopProduct {
	item := &GftShopProduct{
		ShopProductData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
