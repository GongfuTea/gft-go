package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsNavRepo struct {
	*mgo.MgoTreeRepo[*GftCmsNav]
}

var CmsNavRepo = &GftCmsNavRepo{
	mgo.NewMgoTreeRepo[*GftCmsNav]("GftCmsNav"),
}
