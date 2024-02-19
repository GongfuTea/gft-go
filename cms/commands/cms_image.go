package commands

type SaveCmsImage struct {
	Id    string            `json:"id,omitempty"`
	Input SaveCmsImageInput `json:"input"`
}

type DelCmsImage struct {
	Id string `json:"id"`
}

type SaveCmsImageInput struct {
	Name string   `json:"name"`
	Type string   `json:"type"`
	Size int      `json:"size"`
	Url  string   `json:"url"`
	Note string   `json:"note"`
	Tags []string `json:"tags"`
}
