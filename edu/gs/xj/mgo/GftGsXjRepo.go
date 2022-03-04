package mgo

import (
	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/edu/gs/xj"
)

type GftGsXjRepo struct {
	*mgo.MgoRepo
}

var GsXjRepo = &GftGsXjRepo{
	mgo.NewMgoRepo("GftGsXj", xj.NewGftGsXj),
}
