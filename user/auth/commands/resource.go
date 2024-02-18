package commands

import "github.com/GongfuTea/gft-go/user/auth"

type SaveAuthResource struct {
	Input auth.GftAuthResource `json:"input"`
}

type DelAuthResource struct {
	Id string `json:"id"`
}
