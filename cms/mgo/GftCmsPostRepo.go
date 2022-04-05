package mgo

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsPostRepo struct {
	*mgo.MgoRepo[*cms.GftCmsPost]
}

var CmsPostRepo = &GftCmsPostRepo{
	mgo.NewMgoRepo[*cms.GftCmsPost]("GftCmsPost"),
}
