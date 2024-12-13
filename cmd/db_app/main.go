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
		panic(fmt.Errorf("Creating migrate object has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Creating migrate object has successfully")

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Apply migrations (Up)")
		fmt.Println("2. Rollback migrations (Down)")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		switch choice {
		case 1:
			lgr.DebugLogger.Println("Applying migration...")
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				panic(fmt.Errorf("Migration has failed: %s\n", err))
			}
			lgr.InfoLogger.Println("Migration has successfully")
			return
		case 2:
			lgr.DebugLogger.Println("Rolling back migration...")
			if err := m.Down(); err != nil && err != migrate.ErrNoChange {
				panic(fmt.Errorf("Rollback has failed: %s\n", err))
			}
			lgr.InfoLogger.Println("Rollback has successfully")
			return
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
