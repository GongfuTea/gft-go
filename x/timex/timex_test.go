package timex

import (
	"testing"
	"time"
)

func TestTodayStart(t *testing.T) {
	t.Log(TodayStart())
}

func TestTodayEnd(t *testing.T) {
	t.Log(TodayEnd())
}

func TestDateStart(t *testing.T) {
	t.Log(DateStart(time.Now()))
}

func TestDateEnd(t *testing.T) {
	t.Log(DateEnd(time.Now()))
}

func TestParseDateStart(t *testing.T) {
	v, _ := time.ParseInLocation("20060102", "19800917", time.Now().Location())
	t.Log(v)
	// t.Log(v.Local())
	t.Log(DateStart(v))
}
