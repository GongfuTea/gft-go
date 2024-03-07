package commands

type SaveCmsImage struct {
	Id    string            `json:"id,omitempty"`
	Input SaveCmsImageInput `json:"input"`
}

type DelCmsImage struct {
	Id string `json:"id"`
}

type SaveCmsImageInput struct {
	Name string   `bson:"name" json:"name"`
	Type string   `bson:"type" json:"type"`
	Size int      `bson:"size" json:"size"`
	Url  string   `bson:"url" json:"url"`
	Note string   `bson:"note" json:"note"`
	Tags []string `bson:"tags" json:"tags"`
}
