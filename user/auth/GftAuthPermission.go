package auth

type GftAuthPermission struct {
	ResId      string             `bson:"resId" json:"resId"`
	Operations []GftAuthOperation `bson:"operations" json:"operations"`
}
