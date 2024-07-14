package service

import (
	"dota_position_bot/internal/storage"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type BotService interface {
	GetHeroes(*tgbotapi.BotAPI, tgbotapi.Update, string)
	Qq(*tgbotapi.BotAPI, tgbotapi.Update)
	Vsegda(*tgbotapi.BotAPI, tgbotapi.Update)
	Worker(*tgbotapi.BotAPI, tgbotapi.UpdatesChannel)
}

type LogicService struct {
	Repositoriy storage.Storage
}

func NewServiceLogic(store storage.Storage) BotService {
	return &LogicService{Repositoriy: store}
}

func (s *LogicService) GetHeroes(bot *tgbotapi.BotAPI, update tgbotapi.Update, position string) {
	gs, err := s.Repositoriy.GetHeroes(position)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		bot.Send(msg)
	}
	strHeroes := fmt.Sprintf("Герои: %v", gs)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, strHeroes)
	bot.Send(msg)
}

func (s *LogicService) Qq(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello")
	bot.Send(msg)
}

func (s *LogicService) Vsegda(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command")
	bot.Send(msg)
}

func (s *LogicService) Worker(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Command() {
		case "start":
			s.Qq(bot, update)
		case "1pos":
			s.GetHeroes(bot, update, "1pos")
		case "2pos":
			s.GetHeroes(bot, update, "2pos")
		case "3pos":
			s.GetHeroes(bot, update, "3pos")
		case "4pos":
			s.GetHeroes(bot, update, "4pos")
		case "5pos":
			s.GetHeroes(bot, update, "5pos")
		default:
			s.Vsegda(bot, update)
		}
	}
}
