package commands

import "github.com/GongfuTea/gft-go/mall/product"

type SaveProductCategory struct {
	Id    string                      `json:"id,omitempty"`
	Input product.ProductCategoryData `json:"input"`
}

type DelProductCategory struct {
	Id string `json:"id"`
}
