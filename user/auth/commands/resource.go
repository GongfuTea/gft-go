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
	Pid        string                  `bson:"pid,omitempty" json:"pid,omitempty"`
	Code       string                  `bson:"code,omitempty" json:"code,omitempty"`
	Name       string                  `bson:"name,omitempty" json:"name,omitempty"`
	Category   string                  `bson:"category,omitempty" json:"category,omitempty"`
	Operations []auth.GftAuthOperation `bson:"operations" json:"operations"`
	SortOrder  float32                 `bson:"sortOrder" json:"sortOrder"`
}
