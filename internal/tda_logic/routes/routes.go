package routes

import (
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"path/filepath"
)

func NewFiberRouter(userHandler handler.UserHandler) *fiber.App {
	htmlengine := html.New("../../web/templates", ".html")
	app := fiber.New(fiber.Config{
		Views: htmlengine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", nil)
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", nil)
	})
	app.Post("/login", userHandler.LoginUser)
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", nil)
	})
	app.Post("/register", userHandler.InsertUser)
	filepath := filepath.Join("..", "..", "web", "static")
	app.Static("/", filepath)
	return app
}
