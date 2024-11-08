package main

import (
	"flag"
	"fmt"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/routes"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/server"
	"log"
	"net/http"
	"time"

	"github.com/YurcheuskiRadzivon/to_do_app/pkg/db"
	"github.com/YurcheuskiRadzivon/to_do_app/pkg/handlers"
	"github.com/YurcheuskiRadzivon/to_do_app/pkg/routes"
)

func main() {
	var wait time.Duration = 15 * time.Second
	flag.DurationVar(&wait, "timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	fmt.Println("Graceful timeout:", wait)
	app := routes.NewFiberRouter()
	_ = app
	srv := new(server.Server)
	database := db.DatabaseOpen()
	defer database.GetDB().Close()
	var (
		taskService    handlers.TaskHandler
		accountService handlers.UserHandler
	)
	accountService = database
	r := routess.NewMuxRoute(taskService, accountService)
	_ = r
	go func() {

		if err := srv.Run("8080", app); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error occured while running http server", err.Error())
		}
	}()
	srv.SignalWaiting(wait)

}
