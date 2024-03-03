package cms_handlers

import "github.com/GongfuTea/gft-go/core/gql"

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	enagine.AddResolver(&ProductCategoryResolver{})
	enagine.AddResolver(&ShopProductResolver{})
}
