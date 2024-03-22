package base_handlers

import (
	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/base/commands"
	"github.com/GongfuTea/gft-go/base/queries"
)

type DictCategoryResolver struct {
	BaseService *base.BaseService
}

func NewDictCategoryResolver(baseService *base.BaseService) *DictCategoryResolver {
	return &DictCategoryResolver{
		BaseService: baseService,
	}
}

func (r *DictCategoryResolver) SaveDictCategory(cmd commands.SaveDictCategory) (string, error) {
	return r.BaseService.SaveDictCategory(cmd.Id, cmd.Input)
}

func (r *DictCategoryResolver) DictCategories(cmd queries.DictCategories) ([]*base.GftDictCategory, error) {
	return r.BaseService.DictCategoryRepo.All()
}

func (r *DictCategoryResolver) DelDictCategory(cmd commands.DelDictCategory) (bool, error) {
	return r.BaseService.DictCategoryRepo.Del(cmd.Id)
}
