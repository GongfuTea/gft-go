package commands

type SaveCmsCategory struct {
	Id    string               `json:"id,omitempty"`
	Input SaveCmsCategoryInput `json:"input"`
}

type DelCmsCategory struct {
	Id string `json:"id"`
}

type SaveCmsCategoryInput struct {
	Pid       string  `bson:"pid,omitempty" json:"pid,omitempty"`
	Name      string  `bson:"name" json:"name"`
	Code      string  `bson:"code,omitempty" json:"code,omitempty"`
	SortOrder float32 `bson:"sortOrder" json:"sortOrder"`
	Note      string  `bson:"note" json:"note"`
}
