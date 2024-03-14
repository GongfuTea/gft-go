package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsBannerCategoryRepo struct {
	*mgo.MgoTreeRepo[*GftCmsBannerCategory]
}

var CmsBannerCategoryRepo = &GftCmsBannerCategoryRepo{
	mgo.NewMgoTreeRepo[*GftCmsBannerCategory]("GftCmsBannerCategory"),
}
