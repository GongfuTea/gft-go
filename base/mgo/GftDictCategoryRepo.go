package mgo

import (
	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftDictCategoryRepo struct {
	*mgo.MgoTreeRepo
}

var DictCategoryRepo = &GftDictCategoryRepo{
	mgo.NewMgoTreeRepo("GftDictCategory", base.NewGftDictCategory),
}
