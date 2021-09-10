package mgo

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsPostRepo struct {
	*mgo.MgoRepo
}

var CmsPostRepo = &GftCmsPostRepo{
	mgo.NewMgoRepo("GftCmsPost", cms.NewGftCmsPost),
}
