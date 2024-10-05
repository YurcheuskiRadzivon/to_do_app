package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/YurcheuskiRadzivon/to_do_app"
	"github.com/YurcheuskiRadzivon/to_do_app/pkg/handlers"
	"github.com/YurcheuskiRadzivon/to_do_app/pkg/routes"
)

func main() {
	var wait time.Duration = 15 * time.Second
	flag.DurationVar(&wait, "timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	fmt.Println("Graceful timeout:", wait)
	srv := new(to_do_app.Server)
	var taskService handlers.TaskHandler
	
	r := routes.NewMuxRoute(taskService)

	go func() {

		if err := srv.Run("8080", r); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error occured while running http server", err.Error())
		}
	}()
	srv.SignalWaiting(wait)

}
