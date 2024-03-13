package types

import "time"

type ModelBase struct {
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	CreatedBy string    `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	UpdatedBy string    `bson:"updatedBy,omitempty" json:"updatedBy,omitempty"`
}

type TreeModelBase struct {
	Pid   string `bson:"pid,omitempty" json:"pid,omitempty"`
	Code  string `bson:"code" json:"code,omitempty"`
	Mpath string `bson:"mpath" json:"mpath,omitempty"`
}

func (e TreeModelBase) PID() string {
	return e.Pid
}

func (e TreeModelBase) GetCode() string {
	return e.Code
}

func (e TreeModelBase) GetMpath() string {
	return e.Mpath
}

func (e TreeModelBase) HasPid() bool {
	return e.Pid != ""
}

func (e *TreeModelBase) SetMpath(mpath string) {
	e.Mpath = mpath
}
