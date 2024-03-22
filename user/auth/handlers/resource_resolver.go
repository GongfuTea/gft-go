package auth_handlers

import (
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/GongfuTea/gft-go/user/auth/commands"
	"github.com/GongfuTea/gft-go/user/auth/queries"
)

type ResourceResolver struct {
	authService *auth.AuthService
}

func NewResourceResolver(authService *auth.AuthService) *ResourceResolver {
	return &ResourceResolver{
		authService: authService,
	}
}

func (r *ResourceResolver) AuthResources(q queries.AuthResources) ([]auth.GftAuthResource, error) {
	return r.authService.ResourceRepo.All()
}

func (r *ResourceResolver) AuthResource(q queries.AuthResource) (*auth.GftAuthResource, error) {
	return r.authService.ResourceRepo.Get(q.Id)
}

func (r *ResourceResolver) SaveAuthResource(cmd commands.SaveAuthResource) (string, error) {
	return r.authService.SaveResource(cmd.Id, cmd.Input)
}

func (r *ResourceResolver) DelAuthResource(cmd commands.DelAuthResource) (bool, error) {
	return r.authService.ResourceRepo.Del(cmd.Id)
}
