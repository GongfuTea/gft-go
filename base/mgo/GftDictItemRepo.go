package mgo

import (
	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftDictItemRepo struct {
	*mgo.MgoRepo
}

var DictItemRepo = &GftDictItemRepo{
	mgo.NewMgoRepo("GftDictItem", base.NewGftDictItem),
}
