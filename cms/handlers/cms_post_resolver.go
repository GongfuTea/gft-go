package cms_handlers

import (
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/cms"
	"github.com/GongfuTea/gft-go/cms/commands"
	"github.com/GongfuTea/gft-go/cms/queries"
	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type CmsPostResolver struct {
}

func (r *CmsPostResolver) SaveCmsPost(cmd commands.SaveCmsPost) (string, error) {
	if cmd.Id != "" {
		_, err := cms.CmsPostRepo.UpdateById(cmd.Id, cmd.Input)
		if err != nil {
			fmt.Println("err", err)
		}
		return cmd.Id, err
	}

	Post := cms.GftCmsPost{
		Title:       cmd.Input.Title,
		SubTitle:    cmd.Input.SubTitle,
		Abstract:    cmd.Input.Abstract,
		Slug:        cmd.Input.Slug,
		Note:        cmd.Input.Note,
		PublishDate: cmd.Input.PublishDate,
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

func (r *CmsPostResolver) CmsPosts(q queries.CmsPosts) (mgo.QueryPageResult[*cms.GftCmsPost], error) {
	m := bson.M{}
	if q.Filter.Category != "All" {
		if q.Filter.Category == "" {
			m["categoryIds"] = []string{}
		} else {
			m["categoryIds"] = q.Filter.Category
		}
	}
	res, err := cms.CmsPostRepo.Find(m).Page(&q.Filter.PagerFilter)

	return res, err

}

func (r *CmsPostResolver) CmsPost(q queries.CmsPost) (*cms.GftCmsPost, error) {
	return cms.CmsPostRepo.Get(q.Id)
}

func (r *CmsPostResolver) DelCmsPost(cmd commands.DelCmsPost) (bool, error) {
	return cms.CmsPostRepo.Del(cmd.Id)
}
