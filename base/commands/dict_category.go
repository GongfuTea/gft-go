package commands

import "github.com/GongfuTea/gft-go/base"

type SaveDictCategory struct {
	Id    string                `json:"id,omitempty"`
	Input base.DictCategoryData `json:"input"`
}

type DelDictCategory struct {
	Id string `json:"id"`
}
