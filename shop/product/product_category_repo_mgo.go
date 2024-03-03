package product

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftProductCategoryRepo struct {
	*mgo.MgoRepo[*GftProductCategory]
}

var ProductCategoryRepo = &GftProductCategoryRepo{
	mgo.NewMgoRepo[*GftProductCategory]("GftProductCategory"),
}
