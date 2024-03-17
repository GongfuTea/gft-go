package commands

import "github.com/GongfuTea/gft-go/shop/product"

type SaveProductCategory struct {
	Id    string                      `json:"id,omitempty"`
	Input product.ProductCategoryData `json:"input"`
}

type DelProductCategory struct {
	Id string `json:"id"`
}
