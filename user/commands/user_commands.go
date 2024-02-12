package commands

type UserLogin struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}
