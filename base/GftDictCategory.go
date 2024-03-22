package base

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftDictCategory struct {
	types.Entity     `bson:",inline" json:",inline"`
	types.ModelBase  `bson:",inline" json:",inline"`
	DictCategoryData `bson:",inline" json:",inline"`
}

type DictCategoryData struct {
	types.TreeModelBase `bson:",inline" json:",inline"`
	Name                string  `bson:"name" json:"name,omitempty"`
	SortOrder           float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note                string  `bson:"note" json:"note,omitempty"`
}

func NewDictCategory(data DictCategoryData) *GftDictCategory {
	item := &GftDictCategory{
		DictCategoryData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
