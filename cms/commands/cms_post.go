package commands

import (
	"github.com/GongfuTea/gft-go/cms"
)

type SaveCmsPost struct {
	Id    string          `json:"id,omitempty"`
	Input cms.CmsPostData `json:"input"`
}

type DelCmsPost struct {
	Id string `json:"id"`
}
