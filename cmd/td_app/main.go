package main

import (
	"flag"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/config"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/routes"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/server"
	initia "github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/utils/initialization"
	"path/filepath"
	"time"
)

var lgr *logger.Logger = logger.NewLogger()

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			lgr.ErrorLogger.Printf("Caught panic: %v", rec)
		}
	}()
	var wait time.Duration = 10 * time.Second
	flag.DurationVar(&wait, "timeout", time.Second*10, "the duration for which the server gracefully wait for existing connections to finish - e.g. 10s or 1m")
	flag.Parse()
	lgr.InfoLogger.Println("Graceful timeout:", wait, "s.")
	cfgPath := filepath.Join("internal", "config", "config.yaml")
	dsnStr, err := config.GetConfig(cfgPath)
	if err != nil {
		panic(fmt.Errorf("Getting config has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Getting config has successfully")
	userHandler, err := initia.InitializeComponentsUser(dsnStr, lgr)
	if err != nil {
		panic(fmt.Errorf("Initialization user components has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Initialization user components components for router has successfully")
	taskHandler, err := initia.InitializeComponentsTask(dsnStr, lgr)
	if err != nil {
		panic(fmt.Errorf("Initialization task components has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Initialization task components components for router has successfully")
	app := routes.NewFiberRouter(userHandler, taskHandler)
	lgr.InfoLogger.Println("Creating new router has successfully")
	srv := server.NewServer(app)
	lgr.InfoLogger.Println("Creating server object has successfully")
	lgr.DebugLogger.Println("Server running...")
	go func() {

		srv.Run()
	}()
	srv.SignalWaiting(wait, lgr)

}
