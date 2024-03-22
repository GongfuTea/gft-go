package commands

import "github.com/GongfuTea/gft-go/base"

type SaveDictItem struct {
	Id    string            `json:"id,omitempty"`
	Input base.DictItemData `json:"input"`
}

type DelDictItem struct {
	Id string `json:"id"`
}
