package commands

type SaveProductCategory struct {
	Id    string                   `json:"id,omitempty"`
	Input SaveProductCategoryInput `json:"input"`
}

type DelProductCategory struct {
	Id string `json:"id"`
}

type SaveProductCategoryInput struct {
	Pid       string  `json:"pid,omitempty"`
	Name      string  `json:"name"`
	Code      string  `json:"code,omitempty"`
	SortOrder float32 `json:"sortOrder"`
	Note      string  `json:"note"`
}
