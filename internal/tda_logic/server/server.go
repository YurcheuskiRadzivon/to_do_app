package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	app *fiber.App
}

func NewServer(app *fiber.App) *Server {
	return &Server{app: app}
}

func (s *Server) Run() {
	if err := s.app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func (s *Server) SignalWaiting(timeOut time.Duration) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown error:", err)
	} else {
		log.Println("Server shutdown successfully.")
	}

	os.Exit(0)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.Shutdown()
}
