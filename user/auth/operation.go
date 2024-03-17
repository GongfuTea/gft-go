package auth

type GftAuthOperation struct {
	Name string `bson:"name" json:"name"`
	Code string `bson:"code" json:"code"`
}
