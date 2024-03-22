package admin

import (
	"time"

	"github.com/GongfuTea/gft-go/types"
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/google/uuid"
)

type GftAuthAdmin struct {
	types.Entity    `bson:",inline" json:",inline"`
	types.ModelBase `bson:",inline" json:",inline"`
	AuthAdminData   `bson:",inline" json:",inline"`
}

type AuthAdminData struct {
	Name     string             `bson:"name" json:"name"`
	Avatar   string             `bson:"avatar" json:"avatar"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Auths    []auth.GftAuthType `bson:"auths" json:"auths,omitempty"`
	Roles    []auth.GftAuthRole `bson:"roles" json:"roles,omitempty"`
}

func NewAdmin(data AuthAdminData) *GftAuthAdmin {
	item := &GftAuthAdmin{
		AuthAdminData: data,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	return item
}
