package commands

import "github.com/GongfuTea/gft-go/edu/gs/jc"

type SaveGsYxs struct {
	Id    string       `json:"id,omitempty"`
	Input jc.GsYxsData `json:"input"`
}

type DelGsYxs struct {
	Id string `json:"id"`
}
