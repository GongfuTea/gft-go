package base_handlers

import (
	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/base/commands"
	"github.com/GongfuTea/gft-go/base/queries"
)

type DictItemResolver struct {
	BaseService *base.BaseService
}

func NewDictItemResolver(baseService *base.BaseService) *DictItemResolver {
	return &DictItemResolver{
		BaseService: baseService,
	}
}

func (r *DictItemResolver) SaveDictItem(cmd commands.SaveDictItem) (string, error) {
	if cmd.Id != "" {
		_, err := r.BaseService.DictItemRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	item := base.NewDictItem(cmd.Input)
	_, err := r.BaseService.DictItemRepo.Insert(item)
	return item.Id, err
}

func (r *DictItemResolver) DictItems(q queries.DictItems) ([]*base.GftDictItem, error) {
	return r.BaseService.DictItemRepo.FindByCategoryId(q.CategoryId)
}

func (r *DictCategoryResolver) DelDictItem(cmd commands.DelDictItem) (bool, error) {
	return r.BaseService.DictItemRepo.Del(cmd.Id)
}
