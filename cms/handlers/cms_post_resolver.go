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
	cmsService *cms.CmsService
}

func NewPostResolver(cmsService *cms.CmsService) *CmsPostResolver {
	return &CmsPostResolver{
		cmsService: cmsService,
	}
}

func (r *CmsPostResolver) SaveCmsPost(cmd commands.SaveCmsPost) (string, error) {
	if cmd.Id != "" {
		_, err := r.cmsService.PostRepo.UpdateById(cmd.Id, cmd.Input)
		if err != nil {
			fmt.Println("err", err)
		}
		return cmd.Id, err
	}

	Post := cms.NewCmsPost(cmd.Input)
	Post.Id = uuid.NewString()
	Post.CreatedAt = time.Now()
	_, err := r.cmsService.PostRepo.Save(Post)
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
	res, err := r.cmsService.PostRepo.Find(m).Page(&q.Filter.PagerFilter)

	return res, err

}

func (r *CmsPostResolver) CmsPost(q queries.CmsPost) (*cms.GftCmsPost, error) {
	return r.cmsService.PostRepo.Get(q.Id)
}

func (r *CmsPostResolver) DelCmsPost(cmd commands.DelCmsPost) (bool, error) {
	return r.cmsService.PostRepo.Del(cmd.Id)
}
