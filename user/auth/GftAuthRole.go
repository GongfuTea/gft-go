package auth

import (
	"time"
)

type GftAuthRole struct {
	Id          string              `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string              `bson:"name" json:"name"`
	Permissions []GftAuthPermission `bson:"permissions" json:"permissions"`
	CreatedAt   time.Time           `bson:"createdAt" json:"createdAt"`
}
