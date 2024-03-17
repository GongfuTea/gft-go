package commands

import "github.com/GongfuTea/gft-go/user/auth"

type SaveAuthRole struct {
	Id    string            `json:"id,omitempty"`
	Input auth.AuthRoleData `json:"input"`
}

type DelAuthRole struct {
	Id string `json:"id"`
}
