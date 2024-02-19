package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsPostRepo struct {
	*mgo.MgoRepo[*GftCmsPost]
}

var CmsPostRepo = &GftCmsPostRepo{
	mgo.NewMgoRepo[*GftCmsPost]("GftCmsPost"),
}
