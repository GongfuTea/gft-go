package commands

import "github.com/GongfuTea/gft-go/cms"

type SaveCmsCategory struct {
	Id    string              `json:"id,omitempty"`
	Input cms.CmsCategoryData `json:"input"`
}

type DelCmsCategory struct {
	Id string `json:"id"`
}
