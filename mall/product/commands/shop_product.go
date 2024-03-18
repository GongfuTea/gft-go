package commands

import "github.com/GongfuTea/gft-go/mall/product"

type SaveShopProduct struct {
	Id    string                  `json:"id,omitempty"`
	Input product.ShopProductData `json:"input"`
}

type DelShopProduct struct {
	Id string `json:"id"`
}
