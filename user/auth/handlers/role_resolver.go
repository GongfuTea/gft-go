package handlers

import (
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/GongfuTea/gft-go/user/auth/commands"
	"github.com/GongfuTea/gft-go/user/auth/queries"
)

type RoleResolver struct {
}

func (r *RoleResolver) SaveAuthRole(cmd commands.SaveAuthRole) (string, error) {
	if cmd.Id != "" {
		_, err := auth.AuthRoleRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	role := auth.NewAuthRole(cmd.Input)
	_, err := auth.AuthRoleRepo.Save(role)
	return role.Id, err
}

func (r *RoleResolver) AuthRoles(cmd queries.AuthRoles) ([]auth.GftAuthRole, error) {
	return auth.AuthRoleRepo.All()
}

func (r *RoleResolver) DelAuthRole(cmd commands.DelAuthRole) (bool, error) {
	return auth.AuthRoleRepo.Del(cmd.Id)
}
