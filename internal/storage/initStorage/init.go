package initStorage

import (
	"dota_position_bot/internal/config"
	"dota_position_bot/internal/storage"
	"dota_position_bot/internal/storage/PostgreSQL"
	"dota_position_bot/internal/storage/inmemory"
	"errors"
	"log"
)

func InitNewStorage(config *config.Config) (storage.Storage, error) {
	switch config.StorageType() {
	case "inmemory":
		hranilishe, err := inmemory.InitInMemoryStorage()
		if err != nil {
			log.Println("Не удалось создать хранилище")
			return nil, err
		}
		log.Println("Сервис использует хранилище Inmemory")
		return hranilishe, nil
	case "postgres":
		baza, err := PostgreSQL.InitPostgresStorage(config.GetDsn())
		if err != nil {
			log.Println("Не удалось подключиться к БД")
			return nil, err
		}
		log.Println("Сервис использует хранилище PostgreSQL")
		return baza, nil
	default:
		log.Println("Неверно выбран тип хранилища")
	}
	return nil, errors.New("Неверно выбран тип хранилища")
}
