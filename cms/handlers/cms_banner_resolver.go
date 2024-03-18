package cms_handlers

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
)

type CmsBannerResolver struct {
	cmsService *cms.CmsService
}

func NewBannerResolver(cmsService *cms.CmsService) *CmsBannerResolver {
	return &CmsBannerResolver{
		cmsService: cmsService,
	}
}

func (r *CmsBannerResolver) SaveCmsBanner(cmd commands.SaveCmsBanner) (string, error) {
	if cmd.Id != "" {
		_, err := r.cmsService.BannerRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	item := cms.NewCmsBanner(cmd.Input)
	_, err := r.cmsService.BannerRepo.Insert(item)
	return item.Id, err
}

func (r *CmsBannerResolver) CmsBanners(cmd queries.CmsBanners) ([]*cms.GftCmsBanner, error) {
	return r.cmsService.BannerRepo.All()
}

func (r *CmsBannerResolver) CmsBanner(cmd queries.CmsBanner) (*cms.GftCmsBanner, error) {
	return r.cmsService.BannerRepo.Get(cmd.Id)
}

func (r *CmsBannerResolver) DelCmsBanner(cmd commands.DelCmsBanner) (bool, error) {
	return r.cmsService.BannerRepo.Del(cmd.Id)
}
