package base_handlers

import (
	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/core/gql"
)

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	service := base.NewBaseService()
	enagine.AddResolver(NewDictCategoryResolver(service))
	enagine.AddResolver(NewDictItemResolver(service))
}
