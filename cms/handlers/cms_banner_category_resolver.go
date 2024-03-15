package cms_handlers

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
)

type CmsBannerCategoryResolver struct {
}

func NewBannerCategoryResolver() *CmsBannerCategoryResolver {
	return &CmsBannerCategoryResolver{}
}

func (r *CmsBannerCategoryResolver) SaveCmsBannerCategory(cmd commands.SaveCmsBannerCategory) (string, error) {
	if cmd.Id != "" {
		_, err := cms.CmsBannerCategoryRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	category := cms.NewBannerCategory(cmd.Input)
	_, err := cms.CmsBannerCategoryRepo.Save(category)
	return category.Id, err
}

func (r *CmsBannerCategoryResolver) CmsBannerCategories(cmd queries.CmsBannerCategories) ([]*cms.GftCmsBannerCategory, error) {
	return cms.CmsBannerCategoryRepo.All()
}

func (r *CmsBannerCategoryResolver) CmsBannerCategory(q queries.CmsBannerCategory) (*cms.GftCmsBannerCategory, error) {
	return cms.CmsBannerCategoryRepo.Get(q.Id)
}

func (r *CmsBannerCategoryResolver) DelCmsBannerCategory(cmd commands.DelCmsBannerCategory) (bool, error) {
	return cms.CmsBannerCategoryRepo.Del(cmd.Id)
}
