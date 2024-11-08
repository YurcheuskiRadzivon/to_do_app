package user

//переименовать в энтити
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	password string `json:"password"`
}
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

var jwtSecretKey = []byte("very-secret-key")

func GetJwtSecretKey() []byte {
	return jwtSecretKey
}
