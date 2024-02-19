package handlers

import (
	"time"

	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/base/commands"
	"github.com/GongfuTea/gft-go/base/queries"
	"github.com/google/uuid"
)

type DictCategoryResolver struct {
}

func (r *DictCategoryResolver) SaveDictCategory(cmd commands.SaveDictCategory) (string, error) {
	if cmd.Id != "" {
		_, err := base.DictCategoryRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	category := base.GftDictCategory{
		Name:      cmd.Input.Name,
		Note:      cmd.Input.Note,
		SortOrder: cmd.Input.SortOrder,
	}
	category.Id = uuid.NewString()
	category.Pid = cmd.Input.Pid
	category.Code = cmd.Input.Code
	category.CreatedAt = time.Now()
	_, err := base.DictCategoryRepo.Save(&category)
	return category.Id, err
}

func (r *DictCategoryResolver) DictCategories(cmd queries.DictCategories) ([]*base.GftDictCategory, error) {
	return base.DictCategoryRepo.All()
}

func (r *DictCategoryResolver) DelDictCategory(cmd commands.DelDictCategory) (bool, error) {
	return base.DictCategoryRepo.Del(cmd.Id)
}
