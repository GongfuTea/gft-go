package cms

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftCmsBannerCategoryRepo struct {
	*mgo.MgoTreeRepo[*GftCmsBannerCategory]
}

func NewBannerCategoryRepo() *GftCmsBannerCategoryRepo {
	return &GftCmsBannerCategoryRepo{
		mgo.NewMgoTreeRepo[*GftCmsBannerCategory]("GftCmsBannerCategory"),
	}
}
