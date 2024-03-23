package store

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftShopStore struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	ShopStoreData   `bson:",inline" json:",inline"`
}

type ShopStoreData struct {
	Name      string  `bson:"name" json:"name,omitempty"`
	SortOrder float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note      string  `bson:"note" json:"note"`
}

func NewShopStore(data ShopStoreData) *GftShopStore {
	item := &GftShopStore{
		ShopStoreData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
