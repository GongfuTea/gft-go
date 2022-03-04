package mgo

import (
	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/edu/gs/jc"
)

type GftGsYxsRepo struct {
	*mgo.MgoTreeRepo
}

var GsYxsRepo = &GftGsYxsRepo{
	mgo.NewMgoTreeRepo("GftGsYxs", jc.NewGftGsYxs),
}
