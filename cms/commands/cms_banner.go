package commands

import "github.com/GongfuTea/gft-go/cms"

type SaveCmsBanner struct {
	Id    string            `json:"id,omitempty"`
	Input cms.CmsBannerData `json:"input"`
}

type DelCmsBanner struct {
	Id string `json:"id"`
}
