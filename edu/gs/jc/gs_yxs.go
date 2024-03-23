package jc

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftGsYxs struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	GsYxsData       `bson:",inline" json:",inline"`
}

type GsYxsData struct {
	types.TreeModelBase `bson:",inline" json:",inline"`
	Name                string  `bson:"name" json:"name,omitempty"`
	Nickname            string  `bson:"nickname" json:"nickname,omitempty"`
	SortOrder           float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note                string  `bson:"note" json:"note"`
}

func NewGsYxs(data GsYxsData) *GftGsYxs {
	item := &GftGsYxs{
		GsYxsData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
