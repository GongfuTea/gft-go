package commands

import "github.com/GongfuTea/gft-go/edu/gs/jc"

type SaveGsZydm struct {
	Id    string        `json:"id,omitempty"`
	Input jc.GsZydmData `json:"input"`
}

type DelGsZydm struct {
	Id string `json:"id"`
}
