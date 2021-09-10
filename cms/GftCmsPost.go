package cms

import (
	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/types"
)

type GftCmsPost struct {
	*db.DbEntity `bson:",inline"`
	Title        string  `bson:"title" json:"title"`
	SubTitle     string  `bson:"subTitle" json:"subTitle"`
	SortOrder    float32 `bson:"sortOrder" json:"sortOrder,omitempty"`
	Note         string  `bson:"note" json:"note"`
	CreatedBy    string  `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
}

func NewGftCmsPost() types.IEntity {
	return &GftCmsPost{
		DbEntity: db.NewDbEntity(),
	}
}
