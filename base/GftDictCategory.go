package base

import (
	"time"
)

type GftDictCategory struct {
	Id        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Pid       string    `bson:"pid,omitempty" json:"pid,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Slug      string    `bson:"slug" json:"slug"`
	Mpath     string    `bson:"mpath" json:"mpath"`
	SortOrder float32   `bson:"sortOrder" json:"sortOrder"`
	Note      string    `bson:"note" json:"note"`
	CreatedBy string    `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`

	Locale map[string]string `bson:"locale,omitempty" json:"locale,omitempty"`
}
