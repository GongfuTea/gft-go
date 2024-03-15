package cms_handlers

import "github.com/GongfuTea/gft-go/core/gql"

func UseDefaultGqlResolvers() {
	enagine := gql.DefaultSchemaEngine
	enagine.AddResolver(&CmsImageResolver{})
	enagine.AddResolver(&CmsCategoryResolver{})
	enagine.AddResolver(&CmsNavResolver{})
	enagine.AddResolver(&CmsPostResolver{})
	enagine.AddResolver(&CmsBannerResolver{})
	enagine.AddResolver(NewBannerCategoryResolver())
}
