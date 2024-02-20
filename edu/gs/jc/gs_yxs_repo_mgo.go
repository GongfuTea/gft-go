package jc

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftGsYxsRepo struct {
	*mgo.MgoTreeRepo[*GftGsYxs]
}

var GsYxsRepo = &GftGsYxsRepo{
	mgo.NewMgoTreeRepo[*GftGsYxs]("GftGsYxs"),
}
