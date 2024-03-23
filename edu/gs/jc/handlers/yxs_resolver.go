package gs_jc_handlers

import (
	"github.com/GongfuTea/gft-go/edu/gs/jc"
	"github.com/GongfuTea/gft-go/edu/gs/jc/commands"
	"github.com/GongfuTea/gft-go/edu/gs/jc/queries"
)

type YxsResolver struct {
	GsJcService *jc.GsJcService
}

func NewYxsResolver(gsJcService *jc.GsJcService) *YxsResolver {
	return &YxsResolver{
		GsJcService: gsJcService,
	}
}

func (r *YxsResolver) SaveGsYxs(cmd commands.SaveGsYxs) (string, error) {
	return r.GsJcService.SaveGsYxs(cmd.Id, cmd.Input)
}

func (r *YxsResolver) GsYxses(cmd queries.GsYxses) ([]*jc.GftGsYxs, error) {
	return r.GsJcService.YxsRepo.All()
}

func (r *YxsResolver) DelGsYxs(cmd commands.DelGsYxs) (bool, error) {
	return r.GsJcService.YxsRepo.Del(cmd.Id)
}
