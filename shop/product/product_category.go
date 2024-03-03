package product

import (
	"github.com/GongfuTea/gft-go/core/db"
)

type GftProductCategory struct {
	db.DbTreeEntity `bson:",inline" json:",inline"`
	Name            string  `bson:"name" json:"name,omitempty"`
	SortOrder       float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note            string  `bson:"note" json:"note"`
	CreatedBy       string  `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
}
