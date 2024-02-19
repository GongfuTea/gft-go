package handlers

import (
	"time"

	"github.com/GongfuTea/gft-go/base"
	"github.com/GongfuTea/gft-go/base/commands"
	"github.com/GongfuTea/gft-go/base/queries"
	"github.com/google/uuid"
)

type DictItemResolver struct {
}

func (r *DictItemResolver) SaveDictItem(cmd commands.SaveDictItem) (string, error) {
	if cmd.Id != "" {
		_, err := base.DictItemRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	item := base.GftDictItem{
		Name:       cmd.Input.Name,
		CategoryId: cmd.Input.CategoryId,
		Code:       cmd.Input.Code,
		Level:      cmd.Input.Level,
		Nickname:   cmd.Input.Nickname,
		Note:       cmd.Input.Note,
		SortOrder:  cmd.Input.SortOrder,
	}
	item.Id = uuid.NewString()
	item.CreatedAt = time.Now()
	_, err := base.DictItemRepo.Save(&item)
	return item.Id, err
}

func (r *DictItemResolver) DictItems(q queries.DictItems) ([]*base.GftDictItem, error) {
	return base.DictItemRepo.FindByCategoryId(q.CategoryId)
}
