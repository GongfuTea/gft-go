package handlers

import (
	"github.com/GongfuTea/gft-go/user"
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/GongfuTea/gft-go/user/commands"
)

type UserResolver struct {
}

func (r *UserResolver) Login(cmd commands.UserLogin) (*auth.TokenDetails, error) {
	username := cmd.Username
	pass := cmd.Password
	return user.UserRepo.Login(username, pass)
}
