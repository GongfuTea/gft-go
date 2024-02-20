package commands

type SaveGsYxs struct {
	Id    string         `json:"id,omitempty"`
	Input SaveGsYxsInput `json:"input"`
}

type DelGsYxs struct {
	Id string `json:"id"`
}

type SaveGsYxsInput struct {
	Name      string  `json:"name,omitempty"`
	Code      string  `json:"code,omitempty"`
	Pid       string  `json:"pid,omitempty"`
	Nickname  string  `json:"nickname,omitempty"`
	SortOrder float32 `json:"sortOrder,omitempty"`
	Note      string  `json:"note"`
}
