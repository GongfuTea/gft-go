package handlers

import (
	"time"

	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/GongfuTea/gft-go/user/auth/commands"
	"github.com/GongfuTea/gft-go/user/auth/queries"
	"github.com/google/uuid"
)

type RoleResolver struct {
}

func (r *RoleResolver) SaveAuthRole(cmd commands.SaveAuthRole) (string, error) {
	if cmd.Id != "" {
		_, err := auth.AuthRoleRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	role := auth.GftAuthRole{
		Name:        cmd.Input.Name,
		Permissions: cmd.Input.Permissions,
		SortOrder:   cmd.Input.SortOrder,
	}
	role.Id = uuid.NewString()
	role.CreatedAt = time.Now()
	_, err := auth.AuthRoleRepo.Save(&role)
	return role.Id, err
}

func (r *RoleResolver) AuthRoles(cmd queries.AuthRoles) ([]auth.GftAuthRole, error) {
	return auth.AuthRoleRepo.All()
}

func (r *RoleResolver) DelAuthRole(cmd commands.DelAuthRole) (bool, error) {
	return auth.AuthRoleRepo.Del(cmd.Id)
}
