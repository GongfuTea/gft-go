package handlers

import (
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/GongfuTea/gft-go/user/auth/commands"
	"github.com/GongfuTea/gft-go/user/auth/queries"
	"github.com/google/uuid"
)

type RoleResolver struct {
}

func (r *RoleResolver) AddAuthRole(cmd commands.AddAuthRole) (*auth.GftAuthRole, error) {
	role := auth.GftAuthRole{
		Name:        cmd.Input.Name,
		Permissions: cmd.Input.Permissions,
		SortOrder:   cmd.Input.SortOrder,
	}
	role.Id = uuid.NewString()
	return auth.AuthRoleRepo.Save(&role)
}

// func (r *RoleResolver) UpdateAuthRole(cmd commands.UpdateAuthRole) (*auth.GftAuthRole, error) {
// 	role := auth.GftAuthRole{
// 		Name:        cmd.Input.Name,
// 		Permissions: cmd.Input.Permissions,
// 	}
// 	role.Id = uuid.NewString()
// 	return auth.AuthRoleRepo.MgoRepo
// }

func (r *RoleResolver) AuthRoles(cmd queries.AuthRoles) ([]auth.GftAuthRole, error) {
	return auth.AuthRoleRepo.All()
}

func (r *RoleResolver) DelAuthRole(cmd commands.DelAuthRole) (bool, error) {
	return auth.AuthRoleRepo.Del(cmd.Id)
}
