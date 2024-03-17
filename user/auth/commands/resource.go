package commands

import "github.com/GongfuTea/gft-go/user/auth"

type SaveAuthResource struct {
	Id    string                `json:"id,omitempty"`
	Input auth.AuthResourceData `json:"input"`
}

type DelAuthResource struct {
	Id string `json:"id"`
}
