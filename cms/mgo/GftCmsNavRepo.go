package mgo

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsNavRepo struct {
	*mgo.MgoRepo
}

var CmsNavRepo = &GftCmsNavRepo{
	mgo.NewMgoRepo("GftCmsNav", cms.NewGftCmsNav),
}
