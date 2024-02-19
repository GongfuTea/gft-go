package commands

type SaveDictCategory struct {
	Id    string                `json:"id,omitempty"`
	Input SaveDictCategoryInput `json:"input"`
}

type DelDictCategory struct {
	Id string `json:"id"`
}

type SaveDictCategoryInput struct {
	Pid       string  `json:"pid,omitempty"`
	Name      string  `json:"name"`
	Code      string  `json:"code,omitempty"`
	SortOrder float32 `json:"sortOrder"`
	Note      string  `json:"note"`
}
