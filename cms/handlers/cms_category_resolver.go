package cms_handlers

import (
	"time"

	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
	"github.com/google/uuid"
)

type CmsCategoryResolver struct {
}

func (r *CmsCategoryResolver) SaveCmsCategory(cmd commands.SaveCmsCategory) (string, error) {
	if cmd.Id != "" {
		_, err := cms.CmsCategoryRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	category := cms.GftCmsCategory{
		Name:      cmd.Input.Name,
		Note:      cmd.Input.Note,
		SortOrder: cmd.Input.SortOrder,
	}
	category.Id = uuid.NewString()
	category.Pid = cmd.Input.Pid
	category.Code = cmd.Input.Code
	category.CreatedAt = time.Now()
	_, err := cms.CmsCategoryRepo.Save(&category)
	return category.Id, err
}

func (r *CmsCategoryResolver) CmsCategories(cmd queries.CmsCategories) ([]*cms.GftCmsCategory, error) {
	return cms.CmsCategoryRepo.All()
}

func (r *CmsNavResolver) CmsCategory(q queries.CmsCategory) (*cms.GftCmsCategory, error) {
	return cms.CmsCategoryRepo.Get(q.Id)
}

func (r *CmsCategoryResolver) DelCmsCategory(cmd commands.DelCmsCategory) (bool, error) {
	return cms.CmsCategoryRepo.Del(cmd.Id)
}
