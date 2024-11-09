package middleware

import (
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/utils/jwt_service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *fiber.Ctx) error {

	cookie := c.Cookies("tokenAuth")
	if cookie == "" {
		return c.Redirect("/login")
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return jwt_service.GetJwtSecretKey(), nil
	})

	if err != nil || !token.Valid {
		return c.Redirect("/login")
	}

	return c.Next()
}
func RedirectHandler(c *fiber.Ctx) error {
	cookie := c.Cookies("tokenAuth")
	if cookie == "" {
		return c.Redirect("/login")
	}
	return c.Redirect("/tasks")
}
