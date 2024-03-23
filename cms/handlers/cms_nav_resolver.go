package cms_handlers

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
)

type CmsNavResolver struct {
	cmsService *cms.CmsService
}

func NewCmsNavResolver(cmsService *cms.CmsService) *CmsNavResolver {
	return &CmsNavResolver{
		cmsService: cmsService,
	}
}

func (r *CmsNavResolver) SaveCmsNav(cmd commands.SaveCmsNav) (string, error) {
	return r.cmsService.SaveCmsNav(cmd.Id, cmd.Input)
}

func (r *CmsNavResolver) CmsNavs(cmd queries.CmsNavs) ([]*cms.GftCmsNav, error) {
	return r.cmsService.NavRepo.All()
}

func (r *CmsNavResolver) CmsNav(q queries.CmsNav) (*cms.GftCmsNav, error) {
	return r.cmsService.NavRepo.Get(q.Id)
}

func (r *CmsNavResolver) DelCmsNav(cmd commands.DelCmsNav) (bool, error) {
	return r.cmsService.NavRepo.Del(cmd.Id)
}
