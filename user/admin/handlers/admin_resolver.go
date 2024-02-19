package handlers

import (
	"fmt"

	"github.com/GongfuTea/gft-go/user/admin"
	"github.com/GongfuTea/gft-go/user/admin/commands"
	"github.com/GongfuTea/gft-go/user/admin/queries"
	"github.com/GongfuTea/gft-go/user/auth"
)

type AdminResolver struct {
}

func (r *AdminResolver) AdminLogin(cmd commands.AdminLogin) (*auth.TokenDetails, error) {
	user := cmd.Username
	pass := cmd.Password
	fmt.Printf("user: %s, %s", user, pass)

	return admin.AdminRepo.Login(user, pass)
}

func (r *AdminResolver) Admins(q queries.Admins) ([]admin.GftAdmin, error) {
	return admin.AdminRepo.All()
}
