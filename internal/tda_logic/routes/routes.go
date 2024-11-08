package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"path/filepath"
)

func NewFiberRouter() *fiber.App {
	htmlengine := html.New("../../templates", ".html")
	app := fiber.New(fiber.Config{
		Views: htmlengine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("/ main page")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("/login page")
	})
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.SendString("/register register page")
	})

	filepath := filepath.Join("..", "..", "..", "web", "static")
	app.Static("/", filepath)
	return app
}
