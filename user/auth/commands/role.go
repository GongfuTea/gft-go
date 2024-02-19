package commands

import "github.com/GongfuTea/gft-go/user/auth"

type SaveAuthRole struct {
	Id    string            `json:"id,omitempty"`
	Input SaveAuthRoleInput `json:"input"`
}

type DelAuthRole struct {
	Id string `json:"id"`
}

type SaveAuthRoleInput struct {
	Name        string                   `json:"name"`
	Permissions []auth.GftAuthPermission `json:"permissions"`
	SortOrder   float32                  `json:"sortOrder"`
}
