package user_handlers

import (
	"github.com/GongfuTea/gft-go/user"
	"github.com/GongfuTea/gft-go/user/auth"
	"github.com/GongfuTea/gft-go/user/commands"
	"github.com/GongfuTea/gft-go/user/queries"
)

type UserResolver struct {
	userService *user.UserService
}

func NewUserResolver(userService *user.UserService) *UserResolver {
	return &UserResolver{
		userService: userService,
	}
}

func (r *UserResolver) Login(cmd commands.UserLogin) (*auth.TokenDetails, error) {
	username := cmd.Username
	pass := cmd.Password
	return r.userService.UserRepo.Login(username, pass)
}

func (r *UserResolver) Users(q queries.Users) ([]user.GftUser, error) {
	return r.userService.UserRepo.All()
}
