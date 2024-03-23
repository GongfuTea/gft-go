package commands

import "github.com/GongfuTea/gft-go/cms"

type SaveCmsImage struct {
	Id    string           `json:"id,omitempty"`
	Input cms.CmsImageData `json:"input"`
}

type DelCmsImage struct {
	Id string `json:"id"`
}
