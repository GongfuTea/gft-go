package user

import (
	"time"

	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/user/auth"
)

type GftUser struct {
	*db.DbEntity `bson:",inline"`
	Name         string             `bson:"name" json:"name"`
	Avatar       string             `bson:"avatar" json:"avatar"`
	Username     string             `bson:"username" json:"username"`
	Password     string             `bson:"password" json:"password"`
	Auth         []GftUserAuthType  `bson:"auth" json:"auth,omitempty"`
	Roles        []auth.GftAuthRole `bson:"roles" json:"roles,omitempty"`
	CreatedBy    string             `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

func NewGftUser() *GftUser {
	return &GftUser{
		DbEntity: db.NewDbEntity(),
	}
}
