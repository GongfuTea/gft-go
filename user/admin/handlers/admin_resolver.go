package admin_handlers

import (
	"fmt"

	"github.com/GongfuTea/gft-go/user/admin"
	"github.com/GongfuTea/gft-go/user/admin/commands"
	"github.com/GongfuTea/gft-go/user/admin/queries"
	"github.com/GongfuTea/gft-go/user/auth"
)

type AdminResolver struct {
	AdminService *admin.AdminService
}

func NewAdminResolver(adminService *admin.AdminService) *AdminResolver {
	return &AdminResolver{
		AdminService: adminService,
	}
}

func (r *AdminResolver) AdminLogin(cmd commands.AdminLogin) (*auth.TokenDetails, error) {
	user := cmd.Username
	pass := cmd.Password
	fmt.Printf("user: %s, %s", user, pass)

	return r.AdminService.AdminRepo.Login(user, pass)
}

func (r *AdminResolver) Admins(q queries.Admins) ([]*admin.GftAuthAdmin, error) {
	return r.AdminService.AdminRepo.All()
}
