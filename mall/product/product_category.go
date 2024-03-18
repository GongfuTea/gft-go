package product

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftProductCategory struct {
	types.Entity        `bson:",inline" json:",inline"`
	types.ModelBase     `bson:",inline" json:",inline"`
	ProductCategoryData `bson:",inline" json:",inline"`
}

type ProductCategoryData struct {
	types.TreeModelBase `bson:",inline" json:",inline"`
	Name                string  `bson:"name" json:"name,omitempty"`
	SortOrder           float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note                string  `bson:"note" json:"note,omitempty"`
}

func NewProductCategory(data ProductCategoryData) *GftProductCategory {
	item := &GftProductCategory{
		ProductCategoryData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
