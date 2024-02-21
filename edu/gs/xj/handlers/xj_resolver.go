package gs_xj_handlers

import (
	"context"
	"time"

	"github.com/GongfuTea/gft-go/core/mgo"
	"github.com/GongfuTea/gft-go/edu/gs/xj"
	"github.com/GongfuTea/gft-go/edu/gs/xj/commands"
	"github.com/GongfuTea/gft-go/edu/gs/xj/queries"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type XjResolver struct {
}

func (r *XjResolver) SaveGsXj(cmd commands.SaveGsXj) (string, error) {
	if cmd.Id != "" {
		_, err := xj.GsXjRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	zydm := xj.GftGsXj{}
	zydm.Id = uuid.NewString()
	zydm.CreatedAt = time.Now()
	_, err := xj.GsXjRepo.Save(&zydm)
	return zydm.Id, err
}

func (r *XjResolver) GsXjs(q queries.GsXjs) (mgo.QueryPageResult[*xj.GftGsXj], error) {
	return xj.GsXjRepo.Find(context.Background(), bson.M{}).Page(&q.Filter.PagerFilter)
}

func (r *XjResolver) DelGsXj(cmd commands.DelGsXj) (bool, error) {
	return xj.GsXjRepo.Del(cmd.Id)
}
