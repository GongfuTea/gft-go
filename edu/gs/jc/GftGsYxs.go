package jc

import (
	"github.com/GongfuTea/gft-go/core/db"
)

type GftGsYxs struct {
	db.DbTreeEntity `bson:",inline" json:",inline"`
	Name            string  `bson:"name" json:"name,omitempty"`
	Nickname        string  `bson:"nickname" json:"nickname,omitempty"`
	SortOrder       float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note            string  `bson:"note" json:"note"`
	CreatedBy       string  `bson:"createdBy,omitempty" json:"createdBy,omitempty"`

	// Locale map[string]string `bson:"locale,omitempty" json:"locale,omitempty"`
}
