package product

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftShopProductRepo struct {
	*mgo.MgoRepo[*GftShopProduct]
}

func NewShopProductRepo() *GftShopProductRepo {
	return &GftShopProductRepo{
		mgo.NewMgoRepo[*GftShopProduct]("GftShopProduct"),
	}
}
