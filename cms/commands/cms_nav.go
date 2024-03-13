package commands

import "github.com/GongfuTea/gft-go/cms"

type SaveCmsNav struct {
	Id    string          `json:"id,omitempty"`
	Input SaveCmsNavInput `json:"input"`
}

type DelCmsNav struct {
	Id string `json:"id"`
}

type SaveCmsNavInput struct {
	Pid         string          `bson:"pid,omitempty" json:"pid,omitempty"`
	Code        string          `bson:"code,omitempty" json:"code,omitempty"`
	Name        string          `bson:"name,omitempty" json:"name,omitempty"`
	SortOrder   float32         `bson:"sortOrder,omitempty" json:"sortOrder,omitempty"`
	Content     string          `bson:"content" json:"content"`
	State       cms.ActiveState `bson:"state" json:"state"`
	Type        cms.MenuType    `bson:"type" json:"type"`
	TargetIds   []string        `bson:"targetIds" json:"targetIds"`
	AccessLevel cms.AccessLevel `bson:"accessLevel" json:"accessLevel"`
	Note        string          `bson:"note" json:"note"`
	NewWindow   bool            `bson:"newWindow" json:"newWindow"`
}
