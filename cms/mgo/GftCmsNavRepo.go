package mgo

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsNavRepo struct {
	*mgo.MgoRepo[*cms.GftCmsNav]
}

var CmsNavRepo = &GftCmsNavRepo{
	mgo.NewMgoRepo[*cms.GftCmsNav]("GftCmsNav"),
}
