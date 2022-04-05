package mgo

import (
	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftDictItemRepo struct {
	*mgo.MgoRepo[*base.GftDictItem]
}

var DictItemRepo = &GftDictItemRepo{
	mgo.NewMgoRepo[*base.GftDictItem]("GftDictItem"),
}
