package cms_handlers

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/core/gql"
)

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	service := cms.NewCmsService()
	enagine.AddResolver(&CmsImageResolver{})
	enagine.AddResolver(&CmsCategoryResolver{})
	enagine.AddResolver(&CmsNavResolver{})
	enagine.AddResolver(&CmsPostResolver{})
	enagine.AddResolver(NewBannerResolver(service))
	enagine.AddResolver(NewBannerCategoryResolver(service))
}
