package auth_handlers

import (
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/GongfuTea/gft-go/user/auth/commands"
	"github.com/GongfuTea/gft-go/user/auth/queries"
)

type RoleResolver struct {
	authService *auth.AuthService
}

func NewRoleResolver(authService *auth.AuthService) *RoleResolver {
	return &RoleResolver{
		authService: authService,
	}
}

func (r *RoleResolver) SaveAuthRole(cmd commands.SaveAuthRole) (string, error) {
	if cmd.Id != "" {
		_, err := r.authService.RoleRepo.UpdateById(cmd.Id, cmd.Input)
		return cmd.Id, err
	}

	role := auth.NewAuthRole(cmd.Input)
	_, err := r.authService.RoleRepo.Save(role)
	return role.Id, err
}

func (r *RoleResolver) AuthRoles(cmd queries.AuthRoles) ([]auth.GftAuthRole, error) {
	return r.authService.RoleRepo.All()
}

func (r *RoleResolver) DelAuthRole(cmd commands.DelAuthRole) (bool, error) {
	return r.authService.RoleRepo.Del(cmd.Id)
}
