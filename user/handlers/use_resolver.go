package user_handlers

import (
	"github.com/GongfuTea/gft-go/core/gql"
	"github.com/GongfuTea/gft-go/user"
)

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	service := user.NewUserService()
	enagine.AddResolver(NewUserResolver(service))
}
