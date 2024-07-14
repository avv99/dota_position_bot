package app

import (
	"dota_position_bot/internal/config"
	"dota_position_bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func Run(cfg *config.Config, srv service.BotService) error {
	bot, err := tgbotapi.NewBotAPI(cfg.GetToken())
	if err != nil {
		log.Println("Не удалось запустить бота")
		return err
	}

	log.Println("Бот успешно запущен")

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println("Не удалось обновить чат с пользователем")
		return err
	}
	srv.Worker(bot, updates)
	return nil
}
