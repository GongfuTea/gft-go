package cms_handlers

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
)

type CmsBannerCategoryResolver struct {
	cmsService *cms.CmsService
}

func NewBannerCategoryResolver() *CmsBannerCategoryResolver {
	return &CmsBannerCategoryResolver{
		cmsService: cms.NewCmsService(),
	}
}

func (r *CmsBannerCategoryResolver) SaveCmsBannerCategory(cmd commands.SaveCmsBannerCategory) (string, error) {
	return r.cmsService.SaveBannerCategory(cmd.Id, cmd.Input)
}

func (r *CmsBannerCategoryResolver) CmsBannerCategories(cmd queries.CmsBannerCategories) ([]*cms.GftCmsBannerCategory, error) {
	return r.cmsService.BannerCategoryRepo.All()
}

func (r *CmsBannerCategoryResolver) CmsBannerCategory(q queries.CmsBannerCategory) (*cms.GftCmsBannerCategory, error) {
	return r.cmsService.BannerCategoryRepo.Get(q.Id)
}

func (r *CmsBannerCategoryResolver) DelCmsBannerCategory(cmd commands.DelCmsBannerCategory) (bool, error) {
	return r.cmsService.BannerCategoryRepo.Del(cmd.Id)
}
