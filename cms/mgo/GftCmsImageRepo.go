package mgo

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsImageRepo struct {
	*mgo.MgoRepo[*cms.GftCmsImage]
}

var CmsImageRepo = &GftCmsImageRepo{
	mgo.NewMgoRepo[*cms.GftCmsImage]("GftCmsImage"),
}
