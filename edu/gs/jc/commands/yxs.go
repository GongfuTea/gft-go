package commands

type SaveGsYxs struct {
	Id    string         `json:"id,omitempty"`
	Input SaveGsYxsInput `json:"input"`
}

type DelGsYxs struct {
	Id string `json:"id"`
}

type SaveGsYxsInput struct {
	Name      string  `bson:"name,omitempty" json:"name,omitempty"`
	Code      string  `bson:"code,omitempty" json:"code,omitempty"`
	Pid       string  `bson:"pid,omitempty" json:"pid,omitempty"`
	Nickname  string  `bson:"nickname,omitempty" json:"nickname,omitempty"`
	SortOrder float32 `bson:"sortOrder,omitempty" json:"sortOrder,omitempty"`
	Note      string  `bson:"note" json:"note"`
}
