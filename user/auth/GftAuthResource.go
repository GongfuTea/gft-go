package auth

import (
	"time"
)

type GftAuthResource struct {
	Id         string             `bson:"_id,omitempty" json:"id,omitempty"`
	Pid        string             `bson:"pid,omitempty" json:"pid,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Slug       string             `bson:"slug" json:"slug"`
	Operations []GftAuthOperation `bson:"operations" json:"operations"`
	SortOrder  float32            `bson:"sortOrder" json:"sortOrder"`
	CreatedBy  string             `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
}
