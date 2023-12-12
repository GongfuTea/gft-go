package cms

import (
	"github.com/GongfuTea/gft-go/core/db"
)

type GftCmsImage struct {
	db.DbEntity `bson:",inline"`
	Name        string   `bson:"name" json:"name"`
	Url         string   `bson:"url" json:"url"`
	Note        string   `bson:"note" json:"note"`
	Tags        []string `bson:"tags" json:"tags"`
	CreatedBy   string   `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
}
