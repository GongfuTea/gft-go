package user

import (
	"time"
)

type GftUserTimeline struct {
	Id        string    `bson:"_id,omitempty" json:"id,omitempty"`
	UserId    string    `bson:"userId" json:"userId"`
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}
