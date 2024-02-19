package cms_handlers

import (
	"time"

	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
	"github.com/google/uuid"
)

type CmsNavResolver struct {
}

func (r *CmsNavResolver) SaveCmsNav(cmd commands.SaveCmsNav) (string, error) {
	if cmd.Id != "" {
		_, err := cms.CmsNavRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	Nav := cms.GftCmsNav{
		Name:      cmd.Input.Name,
		Note:      cmd.Input.Note,
		SortOrder: cmd.Input.SortOrder,
		Content:   cmd.Input.Content,
		NewWindow: cmd.Input.NewWindow,
		TargetIds: cmd.Input.TargetIds,
	}
	Nav.Id = uuid.NewString()
	Nav.Pid = cmd.Input.Pid
	Nav.Code = cmd.Input.Code
	Nav.CreatedAt = time.Now()
	_, err := cms.CmsNavRepo.Save(&Nav)
	return Nav.Id, err
}

func (r *CmsNavResolver) CmsNavs(cmd queries.CmsNavs) ([]*cms.GftCmsNav, error) {
	return cms.CmsNavRepo.All()
}

func (r *CmsNavResolver) CmsNav(q queries.CmsNav) (*cms.GftCmsNav, error) {
	return cms.CmsNavRepo.Get(q.Id)
}

func (r *CmsNavResolver) DelCmsNav(cmd commands.DelCmsNav) (bool, error) {
	return cms.CmsNavRepo.Del(cmd.Id)
}
