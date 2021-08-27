package user

import (
	"time"
)

type GftUserProfile struct {
	Id string `bson:"_id,omitempty" json:"id,omitempty"`

	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}
