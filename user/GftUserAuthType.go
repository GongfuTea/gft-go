package user

import (
	"time"
)

type GftUserAuthType struct {
	Type      string                 `bson:"type" json:"type"`
	Identity  string                 `bson:"identity" json:"identity"`
	Meta      map[string]interface{} `bson:"meta" json:"meta"`
	CreatedAt time.Time              `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}
