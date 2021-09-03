package auth

import (
	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/types"
)

type GftAuthResource struct {
	*db.DbEntity `bson:",inline"`
	Pid          string             `bson:"pid,omitempty" json:"pid,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Slug         string             `bson:"slug" json:"slug"`
	Operations   []GftAuthOperation `bson:"operations" json:"operations"`
	SortOrder    float32            `bson:"sortOrder" json:"sortOrder"`
	CreatedBy    string             `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
}

func NewGftAuthResource() types.IEntity {
	return &GftAuthResource{
		DbEntity: db.NewDbEntity(),
	}
}
