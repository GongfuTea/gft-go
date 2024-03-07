package commands

type SaveGsZydm struct {
	Id    string          `json:"id,omitempty"`
	Input SaveGsZydmInput `json:"input"`
}

type DelGsZydm struct {
	Id string `json:"id"`
}

type SaveGsZydmInput struct {
	Code      string  `bson:"code" json:"code"`
	Name      string  `bson:"name" json:"name"`
	Level     int     `bson:"level" json:"level"`
	Note      string  `bson:"note" json:"note"`
	Xwlxm     string  `bson:"xwlxm" json:"xwlxm"` // 学位类型 xs/zx
	Xkmlm     string  `bson:"xkmlm" json:"xkmlm"` // 学科门类
	Zscc      string  `bson:"zscc" json:"zscc"`   // 招生层次 s/b/a
	SortOrder float32 `bson:"sortOrder,omitempty" json:"sortOrder,omitempty"`
}
