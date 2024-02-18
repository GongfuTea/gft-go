package handlers

import (
	"fmt"

	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/GongfuTea/gft-go/user/auth/commands"
	"github.com/GongfuTea/gft-go/user/auth/queries"
)

type ResourceResolver struct {
}

func (r *ResourceResolver) SaveAuthResource(cmd commands.SaveAuthResource) (*auth.GftAuthResource, error) {
	res, err := auth.AuthResourceRepo.Save(&cmd.Input)
	if err != nil {
		fmt.Printf("save resource err, %+v", err)
	} else {
		fmt.Printf("save resource, %+v", res)
	}
	return res, err
}

func (r *ResourceResolver) AuthResources(cmd queries.AuthResources) ([]auth.GftAuthResource, error) {
	return auth.AuthResourceRepo.All()
}

func (r *ResourceResolver) DelAuthResource(cmd commands.DelAuthResource) (bool, error) {
	return auth.AuthResourceRepo.Del(cmd.Id)
}
