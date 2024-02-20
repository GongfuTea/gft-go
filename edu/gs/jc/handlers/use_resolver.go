package gs_jc_handlers

import "github.com/GongfuTea/gft-go/core/gql"

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	enagine.AddResolver(&YxsResolver{})
	enagine.AddResolver(&ZydmResolver{})
}
