package main

import (
	"dota_position_bot/internal/app"
	"dota_position_bot/internal/config"
	"dota_position_bot/internal/service"
	"dota_position_bot/internal/storage/initStorage"
	"log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {

		log.Fatal(err)

	}

	stor, err := initStorage.InitNewStorage(cfg)
	if err != nil {

		log.Fatal(err)

	}

	srv := service.NewServiceLogic(stor)

	err = app.Run(cfg, srv)
	if err != nil {
		log.Fatal(err)
	}

}
