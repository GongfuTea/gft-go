package mgo

import (
	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/edu/gs/jc"
)

type GftGsZydmRepo struct {
	*mgo.MgoRepo[*jc.GftGsZydm]
}

var GsZydmRepo = &GftGsZydmRepo{
	mgo.NewMgoRepo("GftGsZydm", jc.NewGftGsZydm),
}
