package cms

import (
	"github.com/GongfuTea/gft-go/core/db"
)

type GftCmsCategory struct {
	*db.DbTreeEntity `bson:",inline"`
	Name             string  `bson:"name" json:"name,omitempty"`
	SortOrder        float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note             string  `bson:"note" json:"note"`
	CreatedBy        string  `bson:"createdBy,omitempty" json:"createdBy,omitempty"`

	Locale map[string]string `bson:"locale,omitempty" json:"locale,omitempty"`
}

func NewGftCmsCategory() *GftCmsCategory {
	return &GftCmsCategory{
		DbTreeEntity: db.NewDbTreeEntity(),
	}
}
