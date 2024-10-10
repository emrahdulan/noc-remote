package main

import (
	"log"

	"bicomsystems.com/network/remote-agent/config"
	"bicomsystems.com/network/remote-agent/db"
	"bicomsystems.com/network/remote-agent/logger"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("FATAL: can't load configuration. err: %s", err)
	}

	logger := logger.Configure(config)

	database, err := db.NewDatabase(config, logger)
	if err != nil {
		logger.Error().Err(err).Str("source", "main.go").Send()
	}
	defer database.Close()

	//repo := repository.NewRepository(database, logger)

	//routes := routes.RegisterRoutes()
}
