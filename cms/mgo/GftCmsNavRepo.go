package mgo

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsNavRepo struct {
	*mgo.MgoTreeRepo[*cms.GftCmsNav]
}

var CmsNavRepo = &GftCmsNavRepo{
	mgo.NewMgoTreeRepo[*cms.GftCmsNav]("GftCmsNav"),
}
