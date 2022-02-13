package base

import (
	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/types"
)

type GftDictItem struct {
	*db.DbEntity `bson:",inline"`
	CategoryId   string  `bson:"categoryId" json:"categoryId"`
	Code         string  `bson:"code" json:"code"`
	Name         string  `bson:"name" json:"name"`
	Nickname     string  `bson:"nickname" json:"nickname"`
	SortOrder    float32 `bson:"sortOrder" json:"sortOrder"`
	Level        int     `bson:"level" json:"level"`
	Note         string  `bson:"note" json:"note"`

	Locale map[string]string `bson:"locale,omitempty" json:"locale,omitempty"`
}

func NewGftDictItem() types.IEntity {
	return &GftDictItem{
		DbEntity: db.NewDbEntity(),
	}
}
