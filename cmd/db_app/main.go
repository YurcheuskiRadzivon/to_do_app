package main

import (
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/config"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var lgr *logger.Logger = logger.NewLogger()

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			lgr.ErrorLogger.Printf("Caught panic: %v", rec)
		}
	}()
	cfgPath := filepath.Join("internal", "config", "config.yaml")
	srcURL := "file://internal//db_logic//migrations"
	dbURL, err := config.GetConfig(cfgPath)
	if err != nil {
		panic(fmt.Errorf("Getting config has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Getting config has successfully")
	m, err := migrate.New(srcURL, dbURL)
	if err != nil {
		panic(fmt.Errorf("Сreating migrate object has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Сreating migrate object has successfully")

	lgr.DebugLogger.Println("Applying migration...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(fmt.Errorf("Migration has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Migration has successfully")

}
