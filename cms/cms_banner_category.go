package cms

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftCmsBannerCategory struct {
	types.Entity          `bson:",inline" json:",inline"`
	types.ModelBase       `bson:",inline" json:",inline"`
	CmsBannerCategoryData `bson:",inline" json:",inline"`
}

type CmsBannerCategoryData struct {
	types.TreeModelBase `bson:",inline" json:",inline"`
	Name                string  `bson:"name" json:"name,omitempty"`
	ImageWidth          int     `bson:"imageWidth" json:"imageWidth,omitempty"`
	ImageHeight         int     `bson:"imageHeight" json:"imageHeight,omitempty"`
	SortOrder           float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note                string  `bson:"note" json:"note,omitempty"`
}

func NewBannerCategory(data CmsBannerCategoryData) *GftCmsBannerCategory {
	item := &GftCmsBannerCategory{
		CmsBannerCategoryData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
