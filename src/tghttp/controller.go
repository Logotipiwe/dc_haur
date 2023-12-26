package tghttp

import (
	"dc_haur/src/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	config "github.com/logotipiwe/dc_go_config_lib"
	"log"
)

func HandleBotUpdates(services *service.Services) {
	handleBotUpdates(services)
}

func CreateTgBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(config.GetConfig("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func handleBotUpdates(services *service.Services) {
	bot := services.TgBotInteractor.GetBot()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			reply, err := services.TgUpdatesHandler.HandleMessageAndReply(update)
			if err != nil {
				println(err.Error())
				reply = services.TgUpdatesHandler.SendUnknownCommandAnswer(update)
			}
			if reply != nil {
				_, err := bot.Send(reply)
				if err != nil {
					println("ERROR WHILE SENDING MESSAGE!")
					println(err)
				}
			}
		}
	}
}
