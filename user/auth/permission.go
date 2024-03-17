package auth

type GftAuthPermission struct {
	ResId  string   `bson:"resId" json:"resId"`
	OptIds []string `bson:"optIds" json:"optIds"`
}
