package jc

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftGsZydmRepo struct {
	*mgo.MgoRepo[*GftGsZydm]
}

var GsZydmRepo = &GftGsZydmRepo{
	mgo.NewMgoRepo[*GftGsZydm]("GftGsZydm"),
}
