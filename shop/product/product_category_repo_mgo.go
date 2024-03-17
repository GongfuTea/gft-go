package product

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftProductCategoryRepo struct {
	*mgo.MgoTreeRepo[*GftProductCategory]
}

func NewProductCategoryRepo() *GftProductCategoryRepo {
	return &GftProductCategoryRepo{
		mgo.NewMgoTreeRepo[*GftProductCategory]("GftProductCategory"),
	}
}
