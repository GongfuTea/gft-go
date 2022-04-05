package mgo

import (
	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftDictCategoryRepo struct {
	*mgo.MgoTreeRepo[*base.GftDictCategory]
}

var DictCategoryRepo = &GftDictCategoryRepo{
	mgo.NewMgoTreeRepo[*base.GftDictCategory]("GftDictCategory"),
}
