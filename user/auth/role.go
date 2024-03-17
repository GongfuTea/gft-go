package auth

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/google/uuid"
)

type GftAuthRole struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	AuthRoleData    `bson:",inline" json:",inline"`
}

type AuthRoleData struct {
	Name        string              `bson:"name" json:"name"`
	Permissions []GftAuthPermission `bson:"permissions" json:"permissions"`
	SortOrder   float32             `bson:"sortOrder" json:"sortOrder"`
	Note        string              `bson:"note" json:"note,omitempty"`
}

func NewAuthRole(data AuthRoleData) *GftAuthRole {
	item := &GftAuthRole{
		AuthRoleData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
