package cms

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftCmsNav struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	CmsNavData      `bson:",inline" json:",inline"`
}

type CmsNavData struct {
	types.TreeModelBase `bson:",inline" json:",inline"`
	Name                string      `bson:"name" json:"name,omitempty"`
	SortOrder           float32     `bson:"sortOrder" json:"sortOrder,omitempty"`
	Content             string      `bson:"content" json:"content"`
	State               ActiveState `bson:"state" json:"state"`
	Type                MenuType    `bson:"type" json:"type"`
	// TypeId          string      `bson:"typeId,omitempty" json:"typeId,omitempty"`
	TargetIds   []string    `bson:"targetIds" json:"targetIds"`
	AccessLevel AccessLevel `bson:"accessLevel" json:"accessLevel"`
	Note        string      `bson:"note" json:"note"`
	NewWindow   bool        `bson:"newWindow" json:"newWindow"`
}

func NewCmsNav(data CmsNavData) *GftCmsNav {
	item := &GftCmsNav{
		CmsNavData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
