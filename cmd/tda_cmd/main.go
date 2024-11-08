package main

import (
	"flag"
	"fmt"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/config"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/routes"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/server"
	initia "github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/utils/initialization"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

func main() {
	var wait time.Duration = 15 * time.Second
	flag.DurationVar(&wait, "timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	fmt.Println("Graceful timeout:", wait)
	cfgPath := filepath.Join("..", "..", "internal", "config", "config.yaml")
	dsnStr, err := config.GetConfig(cfgPath)
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	userHandler, err := initia.InitializeComponentsUser(dsnStr)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	app := routes.NewFiberRouter(userHandler)
	_ = app
	srv := new(server.Server)
	go func() {

		if err := srv.Run("8080", app); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error occured while running http server", err.Error())
		}
	}()
	srv.SignalWaiting(wait)

}
