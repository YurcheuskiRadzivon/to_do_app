package user

type RegUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	password string `json:"password"`
}
