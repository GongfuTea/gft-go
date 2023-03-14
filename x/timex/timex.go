package timex

import "time"

func ParseLocalUTC(t string) time.Time {
	v, _ := time.ParseInLocation("20060102", t, time.Now().Location())
	return v.UTC()
}

func DateStart(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func DateEnd(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())
}

func TodayStart() time.Time {
	return DateStart(time.Now())
}

func TodayEnd() time.Time {
	return DateEnd(time.Now())
}
