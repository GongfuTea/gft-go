package base

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftDictItem struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	DictItemData    `bson:",inline" json:",inline"`
}

type DictItemData struct {
	CategoryId string  `bson:"categoryId" json:"categoryId"`
	Code       string  `bson:"code" json:"code"`
	Name       string  `bson:"name" json:"name"`
	Nickname   string  `bson:"nickname" json:"nickname"`
	SortOrder  float32 `bson:"sortOrder" json:"sortOrder"`
	Level      int     `bson:"level" json:"level"`
	Note       string  `bson:"note" json:"note"`
}

func NewDictItem(data DictItemData) *GftDictItem {
	item := &GftDictItem{
		DictItemData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
