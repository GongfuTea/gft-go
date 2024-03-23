package commands

import "github.com/GongfuTea/gft-go/cms"

type SaveCmsNav struct {
	Id    string         `json:"id,omitempty"`
	Input cms.CmsNavData `json:"input"`
}

type DelCmsNav struct {
	Id string `json:"id"`
}
