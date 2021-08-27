package user

import (
	"time"
)

type GftOrg struct {
	Id        string            `bson:"_id,omitempty" json:"id,omitempty"`
	Pid       string            `bson:"pid,omitempty" json:"pid,omitempty"`
	Name      string            `bson:"name" json:"name"`
	Code      string            `bson:"code" json:"code"`
	Slug      string            `bson:"slug,omitempty" json:"slug,omitempty"`
	Locale    map[string]string `bson:"locale,omitempty" json:"locale,omitempty"`
	CreatedBy string            `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	CreatedAt time.Time         `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}
