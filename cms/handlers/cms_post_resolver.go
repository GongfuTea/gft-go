package cms_handlers

import (
	"time"

	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
	"github.com/google/uuid"
)

type CmsPostResolver struct {
}

func (r *CmsPostResolver) SaveCmsPost(cmd commands.SaveCmsPost) (string, error) {
	if cmd.Id != "" {
		_, err := cms.CmsPostRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	Post := cms.GftCmsPost{
		Title:       cmd.Input.Title,
		SubTitle:    cmd.Input.SubTitle,
		Abstract:    cmd.Input.Abstract,
		Slug:        cmd.Input.Slug,
		Note:        cmd.Input.Note,
		SortOrder:   cmd.Input.SortOrder,
		Content:     cmd.Input.Content,
		NewWindow:   cmd.Input.NewWindow,
		AccessLevel: cmd.Input.AccessLevel,
	}
	Post.Id = uuid.NewString()
	Post.CreatedAt = time.Now()
	_, err := cms.CmsPostRepo.Save(&Post)
	return Post.Id, err
}

func (r *CmsPostResolver) CmsPosts(cmd queries.CmsPosts) ([]*cms.GftCmsPost, error) {
	return cms.CmsPostRepo.All()
}

func (r *CmsPostResolver) CmsPost(q queries.CmsPost) (*cms.GftCmsPost, error) {
	return cms.CmsPostRepo.Get(q.Id)
}

func (r *CmsPostResolver) DelCmsPost(cmd commands.DelCmsPost) (bool, error) {
	return cms.CmsPostRepo.Del(cmd.Id)
}
