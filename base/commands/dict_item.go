package commands

type SaveDictItem struct {
	Id    string            `json:"id,omitempty"`
	Input SaveDictItemInput `json:"input"`
}

type DelDictItem struct {
	Id string `json:"id"`
}

type SaveDictItemInput struct {
	CategoryId string  `bson:"categoryId" json:"categoryId"`
	Code       string  `bson:"code" json:"code"`
	Name       string  `bson:"name" json:"name"`
	Nickname   string  `bson:"nickname" json:"nickname"`
	SortOrder  float32 `bson:"sortOrder" json:"sortOrder"`
	Level      int     `bson:"level" json:"level"`
	Note       string  `bson:"note" json:"note"`
}
