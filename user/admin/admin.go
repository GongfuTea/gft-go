package admin

import (
	"github.com/GongfuTea/gft-go/core/db"
	"github.com/GongfuTea/gft-go/user/auth"
)

type GftAdmin struct {
	db.DbEntity `bson:",inline" json:",inline"`
	Name        string             `bson:"name" json:"name"`
	Avatar      string             `bson:"avatar" json:"avatar"`
	Username    string             `bson:"username" json:"username"`
	Password    string             `bson:"password" json:"password"`
	Auths       []auth.GftAuthType `bson:"auths" json:"auths,omitempty"`
	Roles       []auth.GftAuthRole `bson:"roles" json:"roles,omitempty"`
	CreatedBy   string             `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
}
