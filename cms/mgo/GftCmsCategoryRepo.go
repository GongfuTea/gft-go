package mgo

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsCategoryRepo struct {
	*mgo.MgoTreeRepo
}

var CmsCategoryRepo = &GftCmsCategoryRepo{
	mgo.NewMgoTreeRepo("GftCmsCategory", cms.NewGftCmsCategory),
}
