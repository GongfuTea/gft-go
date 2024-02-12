package commands

type AdminLogin struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}
