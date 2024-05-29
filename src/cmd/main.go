package main

import (
	"go-web-api/src/api"
	"go-web-api/src/config"
	"go-web-api/src/infra/cache"
	"go-web-api/src/infra/persistence/database"
	"go-web-api/src/infra/persistence/migration"
	"go-web-api/src/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migration.Up1()

	api.InitServer(cfg)
}
