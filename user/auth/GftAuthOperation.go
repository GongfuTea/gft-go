package auth

type GftAuthOperation struct {
	Name string `bson:"name" json:"name"`
	Slug string `bson:"slug" json:"slug"`
}
