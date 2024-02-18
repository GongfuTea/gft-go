package auth

import (
	"time"

	"github.com/GongfuTea/gft-go/core/db"
)

type GftAuthRole struct {
	db.DbEntity `bson:",inline" json:",inline"`
	Name        string              `bson:"name" json:"name"`
	Permissions []GftAuthPermission `bson:"permissions" json:"permissions"`
	SortOrder   float32             `bson:"sortOrder" json:"sortOrder"`
	CreatedAt   time.Time           `bson:"createdAt" json:"createdAt"`
}
