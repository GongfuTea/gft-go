package auth

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftAuthResource struct {
	types.Entity     `bson:",inline" json:",inline"`
	types.ModelBase  `bson:",inline" json:",inline"`
	AuthResourceData `bson:",inline" json:",inline"`
}

type AuthResourceData struct {
	types.TreeModelBase `bson:",inline" json:",inline"`
	Name                string             `bson:"name" json:"name,omitempty"`
	Category            string             `bson:"category" json:"category,omitempty"`
	Operations          []GftAuthOperation `bson:"operations" json:"operations"`
	SortOrder           float32            `bson:"sortOrder" json:"sortOrder"`
	Note                string             `bson:"note" json:"note,omitempty"`
}

func NewAuthResource(data AuthResourceData) *GftAuthResource {
	item := &GftAuthResource{
		AuthResourceData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
