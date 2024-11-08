package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpServer *fasthttp.Server
}

func (s *Server) Run(port string, c *fiber.App) error {
	s.httpServer = &fasthttp.Server{
		MaxRequestBodySize: 1 << 20,
		ReadTimeout:        10 * time.Second,
		WriteTimeout:       10 * time.Second,
		Handler:            c.Handler(),
	}
	return s.httpServer.ListenAndServe(":" + port)
}
func (s *Server) SignalWaiting(timeOut time.Duration) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	_ = <-c
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	s.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)

}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown()
}
