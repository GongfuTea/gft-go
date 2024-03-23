package cms_handlers

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
)

type CmsImageResolver struct {
	cmsService *cms.CmsService
}

func NewImageResolver(cmsService *cms.CmsService) *CmsImageResolver {
	return &CmsImageResolver{
		cmsService: cmsService,
	}
}

func (r *CmsImageResolver) SaveCmsImage(cmd commands.SaveCmsImage) (string, error) {
	if cmd.Id != "" {
		_, err := r.cmsService.ImageRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	item := cms.NewCmsImage(cmd.Input)
	_, err := r.cmsService.ImageRepo.Insert(item)
	return item.Id, err
}

func (r *CmsImageResolver) CmsImages(cmd queries.CmsImages) ([]*cms.GftCmsImage, error) {
	return r.cmsService.ImageRepo.All()
}

func (r *CmsImageResolver) DelCmsImage(cmd commands.DelCmsImage) (bool, error) {
	return r.cmsService.ImageRepo.Del(cmd.Id)
}
