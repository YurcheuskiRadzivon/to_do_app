package to_do_app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, r *mux.Router) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        r,
	}
	return s.httpServer.ListenAndServe()
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
	return s.httpServer.Shutdown(ctx)
}
