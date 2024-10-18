package middleware

import (
	"net/http"

	"github.com/YurcheuskiRadzivon/to_do_app/pkg/user"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {

			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return user.GetJwtSecretKey(), nil
		})
		if err != nil || !token.Valid {

			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("jwt")

	if err != nil {

		http.Redirect(w, r, "/login", http.StatusFound)
	} else {

		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
	}
}
