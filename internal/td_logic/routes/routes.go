package routes

import (
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/handler"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"path/filepath"
)

func NewFiberRouter(userHandler handler.UserHandler, taskHandler handler.TaskHandler) *fiber.App {
	htmlengine := html.New("web/templates", ".html")
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
	app.Get("/test", taskHandler.GetTask)
	app.Get("/redirect", middleware.RedirectHandler)
	app.Post("/login", userHandler.LoginUser)
	app.Post("/register", userHandler.InsertUser)
	app.Get("/tasks", middleware.AuthMiddleware, taskHandler.GetTasks)
	app.Get("/tasks/new", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.Render("create_task", nil)
	})
	app.Post("/tasks", middleware.AuthMiddleware, taskHandler.InsertTask)
	app.Get("/tasks/:id", middleware.AuthMiddleware, taskHandler.GetTask)
	app.Put("/tasks/:id", middleware.AuthMiddleware, taskHandler.UpdateTask)
	app.Delete("/tasks/:id", middleware.AuthMiddleware, taskHandler.DeleteTask)
	app.Get("/user", middleware.AuthMiddleware, userHandler.GetUser)
	app.Put("/user", middleware.AuthMiddleware, userHandler.UpdateUser)
	app.Delete("/user", middleware.AuthMiddleware, userHandler.DeleteUser)
	app.Get("/export", middleware.AuthMiddleware, taskHandler.ExportTasks)

	filepath := filepath.Join("web", "static")
	app.Static("/", filepath)
	return app
}
