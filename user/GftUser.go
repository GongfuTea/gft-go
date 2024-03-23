package user

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/google/uuid"
)

type GftUser struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	AuthUserData    `bson:",inline" json:",inline"`
}

type AuthUserData struct {
	Name     string             `bson:"name" json:"name"`
	Avatar   string             `bson:"avatar" json:"avatar"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Auths    []auth.GftAuthType `bson:"auths" json:"auth,omitempty"`
	Roles    []auth.GftAuthRole `bson:"roles" json:"roles,omitempty"`
}

func NewUser(data AuthUserData) *GftUser {
	item := &GftUser{
		AuthUserData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
