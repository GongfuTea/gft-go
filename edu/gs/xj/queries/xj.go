package queries

import (
	"time"

	"github.com/GongfuTea/gft-go/core/db"
)

type GsXjs struct {
	Filter GsXjFilter `json:"filter"`
}

type GsXj struct {
	Id string `json:"id"`
}

type GsXjFilter struct {
	db.PagerFilter
	TimePoint time.Time
}
