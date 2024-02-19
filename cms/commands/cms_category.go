package commands

type SaveCmsCategory struct {
	Id    string               `json:"id,omitempty"`
	Input SaveCmsCategoryInput `json:"input"`
}

type DelCmsCategory struct {
	Id string `json:"id"`
}

type SaveCmsCategoryInput struct {
	Pid       string  `json:"pid,omitempty"`
	Name      string  `json:"name"`
	Code      string  `json:"code,omitempty"`
	SortOrder float32 `json:"sortOrder"`
	Note      string  `json:"note"`
}
