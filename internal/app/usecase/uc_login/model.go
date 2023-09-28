package uc_login

type LoginInfoIn struct {
	UserID   string
	Password string
}

type LoginInfoOut struct {
	Token string `json:"token"`
}
