package commands

type SaveDictCategory struct {
	Id    string                `json:"id,omitempty"`
	Input SaveDictCategoryInput `json:"input"`
}

type DelDictCategory struct {
	Id string `json:"id"`
}

type SaveDictCategoryInput struct {
	Pid       string  `bson:"pid,omitempty" json:"pid,omitempty"`
	Name      string  `bson:"name" json:"name"`
	Code      string  `bson:"code,omitempty" json:"code,omitempty"`
	SortOrder float32 `bson:"sortOrder" json:"sortOrder"`
	Note      string  `bson:"note" json:"note"`
}
