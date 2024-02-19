package commands

import "github.com/GongfuTea/gft-go/user/auth"

type SaveAuthResource struct {
	Id    string                `json:"id,omitempty"`
	Input SaveAuthResourceInput `json:"input"`
}

type DelAuthResource struct {
	Id string `json:"id"`
}

type SaveAuthResourceInput struct {
	Pid        string                  `json:"pid,omitempty"`
	Code       string                  `json:"code,omitempty"`
	Name       string                  `json:"name,omitempty"`
	Category   string                  `json:"category,omitempty"`
	Operations []auth.GftAuthOperation `json:"operations"`
	SortOrder  float32                 `json:"sortOrder"`
}
