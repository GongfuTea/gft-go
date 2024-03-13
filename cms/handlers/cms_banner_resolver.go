package cms_handlers

import (
	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
)

type CmsBannerResolver struct {
}

func (r *CmsBannerResolver) SaveCmsBanner(cmd commands.SaveCmsBanner) (string, error) {
	if cmd.Id != "" {
		_, err := cms.CmsBannerRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	item := cms.NewCmsBanner(cmd.Input)
	_, err := cms.CmsBannerRepo.Insert(item)
	return item.Id, err
}

func (r *CmsBannerResolver) CmsBanners(cmd queries.CmsBanners) ([]*cms.GftCmsBanner, error) {
	return cms.CmsBannerRepo.All()
}

func (r *CmsBannerResolver) CmsBanner(cmd queries.CmsBanner) (*cms.GftCmsBanner, error) {
	return cms.CmsBannerRepo.Get(cmd.Id)
}

func (r *CmsBannerResolver) DelCmsBanner(cmd commands.DelCmsBanner) (bool, error) {
	return cms.CmsBannerRepo.Del(cmd.Id)
}
