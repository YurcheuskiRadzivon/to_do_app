package main

import (
	"github.com/YurcheuskiRadzivon/to_do_app/internal/config"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfgPath := filepath.Join("..", "..", "internal", "config", "config.yaml")
	srcURL := "file://..//..//internal//dba_logic//migrations"
	dbURL, err := config.GetConfig(cfgPath)
	m, err := migrate.New(srcURL, dbURL)
	if err != nil {
		log.Fatal(1, err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(2, err)
	}

}
