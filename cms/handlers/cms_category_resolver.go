package cms_handlers

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
)

type CmsCategoryResolver struct {
	cmsService *cms.CmsService
}

func NewCategoryResolver(cmsService *cms.CmsService) *CmsCategoryResolver {
	return &CmsCategoryResolver{
		cmsService: cmsService,
	}
}

func (r *CmsCategoryResolver) SaveCmsCategory(cmd commands.SaveCmsCategory) (string, error) {
	return r.cmsService.SaveCmsCategory(cmd.Id, cmd.Input)
}

func (r *CmsCategoryResolver) CmsCategories(cmd queries.CmsCategories) ([]*cms.GftCmsCategory, error) {
	return r.cmsService.CategoryRepo.All()
}

func (r *CmsCategoryResolver) CmsCategory(q queries.CmsCategory) (*cms.GftCmsCategory, error) {
	return r.cmsService.CategoryRepo.Get(q.Id)
}

func (r *CmsCategoryResolver) DelCmsCategory(cmd commands.DelCmsCategory) (bool, error) {
	return r.cmsService.CategoryRepo.Del(cmd.Id)
}
