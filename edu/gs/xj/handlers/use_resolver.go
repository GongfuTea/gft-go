package gs_xj_handlers

import "github.com/GongfuTea/gft-go/core/gql"

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	enagine.AddResolver(&XjResolver{})
}
