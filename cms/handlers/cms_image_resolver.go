package cms_handlers

import (
	"time"

	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
	"github.com/google/uuid"
)

type CmsImageResolver struct {
}

func (r *CmsImageResolver) SaveCmsImage(cmd commands.SaveCmsImage) (string, error) {
	if cmd.Id != "" {
		_, err := cms.CmsImageRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	Image := cms.GftCmsImage{
		Name: cmd.Input.Name,
		Note: cmd.Input.Note,
		Type: cmd.Input.Type,
		Size: cmd.Input.Size,
		Url:  cmd.Input.Url,
		Tags: cmd.Input.Tags,
	}
	Image.Id = uuid.NewString()
	Image.CreatedAt = time.Now()
	_, err := cms.CmsImageRepo.Save(&Image)
	return Image.Id, err
}

func (r *CmsImageResolver) CmsImages(cmd queries.CmsImages) ([]*cms.GftCmsImage, error) {
	return cms.CmsImageRepo.All()
}

func (r *CmsImageResolver) DelCmsImage(cmd commands.DelCmsImage) (bool, error) {
	return cms.CmsImageRepo.Del(cmd.Id)
}
