package db

type PagerFilter struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}
