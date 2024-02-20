package gs_jc_handlers

import (
	"time"

	"github.com/GongfuTea/gft-go/edu/gs/jc"
	"github.com/GongfuTea/gft-go/edu/gs/jc/commands"
	"github.com/GongfuTea/gft-go/edu/gs/jc/queries"
	"github.com/google/uuid"
)

type YxsResolver struct {
}

func (r *YxsResolver) SaveGsYxs(cmd commands.SaveGsYxs) (string, error) {
	if cmd.Id != "" {
		_, err := jc.GsYxsRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	yxs := jc.GftGsYxs{
		Name:      cmd.Input.Name,
		Nickname:  cmd.Input.Nickname,
		Note:      cmd.Input.Note,
		SortOrder: cmd.Input.SortOrder,
	}
	yxs.Id = uuid.NewString()
	yxs.Code = cmd.Input.Code
	yxs.Pid = cmd.Input.Pid
	yxs.CreatedAt = time.Now()
	_, err := jc.GsYxsRepo.Save(&yxs)
	return yxs.Id, err
}

func (r *YxsResolver) GsYxses(cmd queries.GsYxses) ([]*jc.GftGsYxs, error) {
	return jc.GsYxsRepo.All()
}

func (r *YxsResolver) DelGsYxs(cmd commands.DelGsYxs) (bool, error) {
	return jc.GsYxsRepo.Del(cmd.Id)
}
