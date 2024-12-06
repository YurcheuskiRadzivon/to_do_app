package server

import (
	"context"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/gofiber/fiber/v2"
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
		if err != nil {
			panic(fmt.Errorf("Starting server has failed: %s\n", err))
		}

	}
}

func (s *Server) SignalWaiting(timeOut time.Duration, lgr *logger.Logger) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		panic(fmt.Errorf("Server Shutdown has failed: %s\n", err))

	} else {
		lgr.InfoLogger.Println("Server shutdown successfully.")
	}

	os.Exit(0)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.Shutdown()
}
