package gs_jc_handlers

import (
	"github.com/GongfuTea/gft-go/edu/gs/jc"
	"github.com/GongfuTea/gft-go/edu/gs/jc/commands"
	"github.com/GongfuTea/gft-go/edu/gs/jc/queries"
)

type ZydmResolver struct {
	GsJcService *jc.GsJcService
}

func NewZydmResolver(gsJcService *jc.GsJcService) *YxsResolver {
	return &YxsResolver{
		GsJcService: gsJcService,
	}
}

func (r *YxsResolver) SaveGsZydm(cmd commands.SaveGsZydm) (string, error) {
	if cmd.Id != "" {
		_, err := r.GsJcService.ZydmRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	item := jc.NewGsZydm(cmd.Input)
	_, err := r.GsJcService.ZydmRepo.Insert(item)
	return item.Id, err
}

func (r *ZydmResolver) GsZydms(cmd queries.GsZydms) ([]*jc.GftGsZydm, error) {
	return r.GsJcService.ZydmRepo.All()
}

func (r *ZydmResolver) DelGsZydm(cmd commands.DelGsZydm) (bool, error) {
	return r.GsJcService.ZydmRepo.Del(cmd.Id)
}
