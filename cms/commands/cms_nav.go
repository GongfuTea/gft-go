package commands

import "github.com/GongfuTea/gft-go/cms"

type SaveCmsNav struct {
	Id    string          `json:"id,omitempty"`
	Input SaveCmsNavInput `json:"input"`
}

type DelCmsNav struct {
	Id string `json:"id"`
}

type SaveCmsNavInput struct {
	Pid         string          `json:"pid,omitempty"`
	Code        string          `json:"code,omitempty"`
	Name        string          `json:"name,omitempty"`
	SortOrder   float32         `json:"sortOrder,omitempty"`
	Content     string          `json:"content"`
	State       cms.PostState   `json:"state"`
	Type        cms.MenuType    `json:"type"`
	TargetIds   []string        `json:"targetIds"`
	AccessLevel cms.AccessLevel `json:"accessLevel"`
	Note        string          `json:"note"`
	NewWindow   bool            `json:"newWindow"`
}
