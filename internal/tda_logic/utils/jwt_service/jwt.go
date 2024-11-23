package jwt_service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
)

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
func GetUserId(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecretKey(), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["sub_id"].(float64))
		log.Println(userID)
		return userID, nil
	}
	return 0, fmt.Errorf("invalid token")
}
func GetUserNickname(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecretKey(), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		nickname := claims["name"].(string)
		log.Println(nickname)
		return nickname, nil
	}
	return "", fmt.Errorf("invalid token")
}
func GetEmailNickname(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecretKey(), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["name"].(string)
		log.Println(email)
		return email, nil
	}
	return "", fmt.Errorf("invalid token")
}
