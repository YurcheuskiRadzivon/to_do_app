package model

type User struct {
	ID       int    `json:"id,omitempy"`
	Name     string `json:"name,omitempy"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserHash struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
