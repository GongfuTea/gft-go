package cms

import (
	"github.com/GongfuTea/gft-go/core/db"
)

type GftCmsNav struct {
	db.DbTreeEntity `bson:",inline"`
	Name            string      `bson:"name" json:"name,omitempty"`
	SortOrder       float32     `bson:"sortOrder" json:"sortOrder,omitempty"`
	Content         string      `bson:"content" json:"content"`
	State           PostState   `bson:"state" json:"state"`
	Type            MenuType    `bson:"type" json:"type"`
	AccessLevel     AccessLevel `bson:"accessLevel" json:"accessLevel"`
	TypeId          string      `bson:"typeId,omitempty" json:"typeId,omitempty"`
	Note            string      `bson:"note" json:"note"`
	NewWindow       bool        `bson:"newWindow" json:"newWindow"`
}
