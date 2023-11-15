package xlsx

import "strconv"

type ItemMap map[string]string

func (it ItemMap) GetField(field string) string {
	if val, ok := it[field]; ok {
		return val
	}
	return ""
}

func (it ItemMap) GetFieldDefault(field string, defaults string) string {
	if val, ok := it[field]; ok && val != "" {
		return val
	}
	return defaults
}

func (it ItemMap) GetFloat(field string) float64 {
	str := it.GetField(field)
	v, _ := strconv.ParseFloat(str, 32)
	return v
}

func (it ItemMap) GetInt(field string) int {
	str := it.GetField(field)
	v, _ := strconv.Atoi(str)
	return v
}
