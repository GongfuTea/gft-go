package mgo

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsCategoryRepo struct {
	*mgo.MgoTreeRepo[*cms.GftCmsCategory]
}

var CmsCategoryRepo = &GftCmsCategoryRepo{
	mgo.NewMgoTreeRepo[*cms.GftCmsCategory]("GftCmsCategory"),
}
