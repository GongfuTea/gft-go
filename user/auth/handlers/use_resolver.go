package auth_handlers

import (
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/user/auth"
)

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	service := auth.NewAuthService()
	enagine.AddResolver(NewRoleResolver(service))
	enagine.AddResolver(NewResourceResolver(service))
}
