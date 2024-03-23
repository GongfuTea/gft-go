package cms

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftCmsCategory struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	CmsCategoryData `bson:",inline" json:",inline"`
}

type CmsCategoryData struct {
	types.TreeModelBase `bson:",inline" json:",inline"`
	Name                string  `bson:"name" json:"name,omitempty"`
	SortOrder           float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note                string  `bson:"note" json:"note"`
}

func NewCmsCategory(data CmsCategoryData) *GftCmsCategory {
	item := &GftCmsCategory{
		CmsCategoryData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
