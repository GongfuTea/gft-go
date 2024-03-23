package gs_jc_handlers

import (
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/edu/gs/jc"
)

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	service := jc.NewGsJcService()
	enagine.AddResolver(NewYxsResolver(service))
	enagine.AddResolver(NewZydmResolver(service))
}
