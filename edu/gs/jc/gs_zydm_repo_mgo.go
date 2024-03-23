package jc

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftGsZydmRepo struct {
	*mgo.MgoRepo[*GftGsZydm]
}

func NewGsZydmRepo() *GftGsZydmRepo {
	return &GftGsZydmRepo{
		mgo.NewMgoRepo[*GftGsZydm]("GftGsZydm"),
	}
}
