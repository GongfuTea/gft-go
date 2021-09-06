package mgo

import (
	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftDictRepo struct {
	*mgo.MgoRepo
}

var DictRepo = &GftDictRepo{
	mgo.NewMgoRepo("GftDict", base.NewGftDict),
}
