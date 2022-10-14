package mgo

import (
	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/edu/gs/xg/jz/gjjxj"
)

type GftXgJzGjjxjRepo struct {
	*mgo.MgoRepo[*gjjxj.GftXgJzGjjxj]
}

var XgJzGjjxjRepo = &GftXgJzGjjxjRepo{
	mgo.NewMgoRepo[*gjjxj.GftXgJzGjjxj]("GftXgJzGjjxj"),
}
