package gs_jc_handlers

import (
	"time"

	"github.com/GongfuTea/gft-go/edu/gs/jc"
	"github.com/GongfuTea/gft-go/edu/gs/jc/commands"
	"github.com/GongfuTea/gft-go/edu/gs/jc/queries"
	"github.com/google/uuid"
)

type ZydmResolver struct {
}

func (r *YxsResolver) SaveGsZydm(cmd commands.SaveGsZydm) (string, error) {
	if cmd.Id != "" {
		_, err := jc.GsZydmRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	zydm := jc.GftGsZydm{
		Name:  cmd.Input.Name,
		Code:  cmd.Input.Code,
		Level: cmd.Input.Level,
		Note:  cmd.Input.Note,
		Xwlxm: cmd.Input.Xwlxm,
		Xkmlm: cmd.Input.Xkmlm,
		Zscc:  cmd.Input.Zscc,
	}
	zydm.Id = uuid.NewString()
	zydm.CreatedAt = time.Now()
	_, err := jc.GsZydmRepo.Save(&zydm)
	return zydm.Id, err
}

func (r *ZydmResolver) GsZydms(cmd queries.GsZydms) ([]*jc.GftGsZydm, error) {
	return jc.GsZydmRepo.All()
}

func (r *ZydmResolver) DelGsZydm(cmd commands.DelGsZydm) (bool, error) {
	return jc.GsZydmRepo.Del(cmd.Id)
}
