package mgo

import (
	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/edu/gs/zs"
)

type GftGsZsLqkRepo struct {
	*mgo.MgoRepo[*zs.GftGsZsLqk]
}

var GsZsLqkRepo = &GftGsZsLqkRepo{
	mgo.NewMgoRepo[*zs.GftGsZsLqk]("GftGsZsLqk"),
}
