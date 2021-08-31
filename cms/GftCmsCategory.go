package cms

import (
	"time"
)

type GftCmsCategory struct {
	Id        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name,omitempty"`
	Slug      string    `bson:"slug" json:"slug,omitempty"`
	Pid       string    `bson:"pid,omitempty" json:"pid,omitempty"`
	Mpath     string    `bson:"mpath" json:"mpath,omitempty"`
	SortOrder float32   `bson:"sortOrder" json:"sortOrder,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}
