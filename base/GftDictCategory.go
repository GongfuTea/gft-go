package base

import (
	"github.com/GongfuTea/gft-go/core/db"
)

type GftDictCategory struct {
	*db.DbTreeEntity `bson:",inline"`
	Name             string  `bson:"name" json:"name"`
	SortOrder        float32 `bson:"sortOrder" json:"sortOrder"`
	Note             string  `bson:"note" json:"note"`
	CreatedBy        string  `bson:"createdBy,omitempty" json:"createdBy,omitempty"`

	Locale map[string]string `bson:"locale,omitempty" json:"locale,omitempty"`
}

 