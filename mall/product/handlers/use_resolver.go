package cms_handlers

import (
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/mall/product"
)

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	service := product.NewProductService()
	enagine.AddResolver(NewProductCategoryResolver(service))
	enagine.AddResolver(NewShopProductResolver(service))
}
