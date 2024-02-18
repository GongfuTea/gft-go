package commands

import "github.com/GongfuTea/gft-go/user/auth"

type AddAuthRole struct {
	Input AddAuthRoleInput `json:"input"`
}

type UpdateAuthRole struct {
	Input UpdateAuthRoleInput `json:"input"`
}

type DelAuthRole struct {
	Id string `json:"id"`
}

type AddAuthRoleInput struct {
	Name        string                   `json:"name"`
	Permissions []auth.GftAuthPermission `json:"permissions"`
	SortOrder   float32                  `json:"sortOrder"`
}

type UpdateAuthRoleInput struct {
	Id          string                   `json:"id"`
	Name        string                   `json:"name"`
	Permissions []auth.GftAuthPermission `json:"permissions"`
	SortOrder   float32                  `json:"sortOrder"`
}
