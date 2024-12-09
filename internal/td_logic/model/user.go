package model

type User struct {
	ID       int    `json:"id,omitempy"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password,omitempy"`
}
type UserHash struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
