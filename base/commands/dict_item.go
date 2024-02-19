package commands

type SaveDictItem struct {
	Id    string            `json:"id,omitempty"`
	Input SaveDictItemInput `json:"input"`
}

type DelDictItem struct {
	Id string `json:"id"`
}

type SaveDictItemInput struct {
	CategoryId string  `json:"categoryId"`
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Nickname   string  `json:"nickname"`
	SortOrder  float32 `json:"sortOrder"`
	Level      int     `json:"level"`
	Note       string  `json:"note"`
}
