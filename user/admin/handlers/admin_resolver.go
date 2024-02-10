package handlers

import (
	"github.com/GongfuTea/gft-go/user/admin"
	"github.com/GongfuTea/gft-go/user/auth"
)

type AdminResolver struct {
}

func (r *AdminResolver) AdminLogin(cmd AdminLoginCmd) (*auth.TokenDetails, error) {
	user := cmd.Username
	pass := cmd.Password
	return admin.AdminRepo.Login(user, pass)
}

func (r *AdminResolver) AdminLogin2(cmd AdminLoginQuery) (*AdminLoginQueryResp, error) {
	return nil, nil
}

type AdminLoginCmd struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type AdminLoginQuery struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type AdminLoginQueryResp struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}
