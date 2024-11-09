package jwt_service

import "github.com/golang-jwt/jwt"

var jwtSecretKey = []byte("very-secret-key")

func GetJwtSecretKey() []byte {
	return jwtSecretKey
}

func CreateToken(payload jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString(GetJwtSecretKey())
	if err != nil {
		return "", err
	}
	return t, nil
}
