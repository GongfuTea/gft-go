package commands

type SaveGsZydm struct {
	Id    string          `json:"id,omitempty"`
	Input SaveGsZydmInput `json:"input"`
}

type DelGsZydm struct {
	Id string `json:"id"`
}

type SaveGsZydmInput struct {
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Level     int     `json:"level"`
	Note      string  `json:"note"`
	Xwlxm     string  `json:"xwlxm"` // 学位类型 xs/zx
	Xkmlm     string  `json:"xkmlm"` // 学科门类
	Zscc      string  `json:"zscc"`  // 招生层次 s/b/a
	SortOrder float32 `json:"sortOrder,omitempty"`
}
