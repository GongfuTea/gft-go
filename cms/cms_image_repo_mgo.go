package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsImageRepo struct {
	*mgo.MgoRepo[*GftCmsImage]
}

var CmsImageRepo = &GftCmsImageRepo{
	mgo.NewMgoRepo[*GftCmsImage]("GftCmsImage"),
}
