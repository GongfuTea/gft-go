package handlers

import (
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
	return admin.AdminRepo.Login(user, pass)
}

func (r *AdminResolver) AdminLogin2(cmd queries.AdminLoginQuery) (*AdminLoginQueryResp, error) {
	return nil, nil
}

func (r *AdminResolver) AdminLogin3(cmd queries.AdminLogin2Query) (AdminLoginQueryResp, error) {
	return AdminLoginQueryResp{}, nil
}

func (r *AdminResolver) AdminLogin4(cmd queries.AdminLogin3Query) (bool, error) {
	return true, nil
}

func (r *AdminResolver) AdminLogin5(cmd queries.AdminLogin4Query) (bool, error) {
	return true, nil
}

type AdminLoginQueryResp struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}
