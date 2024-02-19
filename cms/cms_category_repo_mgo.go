package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsCategoryRepo struct {
	*mgo.MgoTreeRepo[*GftCmsCategory]
}

var CmsCategoryRepo = &GftCmsCategoryRepo{
	mgo.NewMgoTreeRepo[*GftCmsCategory]("GftCmsCategory"),
}
