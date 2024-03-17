package handlers

import (
	"time"

	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/GongfuTea/gft-go/user/auth/commands"
	"github.com/GongfuTea/gft-go/user/auth/queries"
	"github.com/google/uuid"
)

type ResourceResolver struct {
}

func (r *ResourceResolver) AuthResources(q queries.AuthResources) ([]auth.GftAuthResource, error) {
	return auth.AuthResourceRepo.All()
}

func (r *ResourceResolver) AuthResource(q queries.AuthResource) (*auth.GftAuthResource, error) {
	return auth.AuthResourceRepo.Get(q.Id)
}

func (r *ResourceResolver) SaveAuthResource(cmd commands.SaveAuthResource) (string, error) {
	if cmd.Id != "" {
		_, err := auth.AuthResourceRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	res := auth.NewAuthResource(cmd.Input)
	res.Id = uuid.NewString()
	res.CreatedAt = time.Now()
	_, err := auth.AuthResourceRepo.Save(res)
	return res.Id, err
}

func (r *ResourceResolver) DelAuthResource(cmd commands.DelAuthResource) (bool, error) {
	return auth.AuthResourceRepo.Del(cmd.Id)
}
