package commands

import "github.com/GongfuTea/gft-go/cms"

type SaveCmsBannerCategory struct {
	Id    string                    `json:"id,omitempty"`
	Input cms.CmsBannerCategoryData `json:"input"`
}

type DelCmsBannerCategory struct {
	Id string `json:"id"`
}
