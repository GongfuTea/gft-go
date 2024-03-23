package cms_handlers

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/gql"
)

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	service := cms.NewCmsService()
	enagine.AddResolver(NewImageResolver(service))
	enagine.AddResolver(NewCategoryResolver(service))
	enagine.AddResolver(NewCmsNavResolver(service))
	enagine.AddResolver(NewPostResolver(service))
	enagine.AddResolver(NewBannerResolver(service))
	enagine.AddResolver(NewBannerCategoryResolver(service))
}
