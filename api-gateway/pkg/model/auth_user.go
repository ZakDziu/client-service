package model

type AuthUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
