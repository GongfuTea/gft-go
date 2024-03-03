package product

import (
	"github.com/GongfuTea/gft-go/core/mgo"
)

type GftShopProductRepo struct {
	*mgo.MgoRepo[*GftShopProduct]
}

var ShopProductRepo = &GftShopProductRepo{
	mgo.NewMgoRepo[*GftShopProduct]("GftShopProduct"),
}
