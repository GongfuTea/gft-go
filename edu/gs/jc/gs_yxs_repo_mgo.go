package jc

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftGsYxsRepo struct {
	*mgo.MgoTreeRepo[*GftGsYxs]
}

func NewGsYxsRepo() *GftGsYxsRepo {
	return &GftGsYxsRepo{
		mgo.NewMgoTreeRepo[*GftGsYxs]("GftGsYxs"),
	}
}
