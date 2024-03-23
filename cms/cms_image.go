package cms

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftCmsImage struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	CmsImageData    `bson:",inline" json:",inline"`
}

type CmsImageData struct {
	Name string   `bson:"name" json:"name"`
	Type string   `bson:"type" json:"type"`
	Size int      `bson:"size" json:"size"`
	Url  string   `bson:"url" json:"url"`
	Note string   `bson:"note" json:"note"`
	Tags []string `bson:"tags" json:"tags"`
}

func NewCmsImage(data CmsImageData) *GftCmsImage {
	item := &GftCmsImage{
		CmsImageData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
