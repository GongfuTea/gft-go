package commands

type SaveProductCategory struct {
	Id    string                   `json:"id,omitempty"`
	Input SaveProductCategoryInput `json:"input"`
}

type DelProductCategory struct {
	Id string `json:"id"`
}

type SaveProductCategoryInput struct {
	Pid       string  `json:"pid,omitempty" bson:"pid,omitempty"`
	Name      string  `json:"name" bson:"name"`
	Code      string  `json:"code,omitempty" bson:"code,omitempty"`
	SortOrder float32 `json:"sortOrder" bson:"sortOrder"`
	Note      string  `json:"note" bson:"note"`
}
