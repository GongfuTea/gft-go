package auth

import (
	"github.com/GongfuTea/gft-go/core/db"
)

type GftAuthResource struct {
	db.DbTreeEntity `bson:",inline"`
	Name            string             `bson:"name" json:"name,omitempty"`
	Category        string             `bson:"category" json:"category,omitempty"`
	Operations      []GftAuthOperation `bson:"operations" json:"operations"`
	SortOrder       float32            `bson:"sortOrder" json:"sortOrder"`
	CreatedBy       string             `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
}

// func NewGftAuthResource() *GftAuthResource {
// 	return &GftAuthResource{
// 		DbTreeEntity: db.NewDbTreeEntity(),
// 	}
// }
