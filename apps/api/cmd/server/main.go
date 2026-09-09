package main

import (
	"log"

	"github.com/codeatlas/api/internal/config"
	"github.com/codeatlas/api/internal/platform/db"
	"github.com/codeatlas/api/internal/platform/httpserver"
)

func main() {
	cfg := config.Load()

	database := db.Connect(cfg.DatabaseURL)
	db.Migrate(database)

	router := httpserver.NewRouter(database, cfg)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
