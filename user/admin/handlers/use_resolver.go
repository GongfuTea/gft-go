package admin_handlers

import (
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/user/admin"
)

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	service := admin.NewAdminService()
	enagine.AddResolver(NewAdminResolver(service))
}
