package routes

import (
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/handler"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"path/filepath"
)

func NewFiberRouter(userHandler handler.UserHandler, taskHandler handler.TaskHandler) *fiber.App {
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
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", nil)
	})
	app.Get("/test", taskHandler.GetTasks)
	app.Get("/redirect", middleware.RedirectHandler)
	app.Post("/login", userHandler.LoginUser)
	app.Post("/register", userHandler.InsertUser)
	app.Get("/tasks", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.Render("tasks", nil)
	})
	filepath := filepath.Join("..", "..", "web", "static")
	app.Static("/", filepath)
	return app
}
