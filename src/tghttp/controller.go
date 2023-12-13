package tghttp

import (
	"database/sql"
	handler "dc_haur/src/internal"
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/repo"
	"dc_haur/src/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	config "github.com/logotipiwe/dc_go_config_lib"
	"log"
)

func StartTgBot(db *sql.DB) {
	tgBot(db)
}

func tgBot(db *sql.DB) {
	tgbot, err := tgbotapi.NewBotAPI(config.GetConfig("BOT_TOKEN"))
	bot := domain.Bot{Instance: tgbot}
	if err != nil {
		log.Panic(err)
	}
	bot.Instance.Debug = true
	log.Printf("Authorized on account %s", bot.Instance.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.Instance.GetUpdatesChan(u)

	repos := repo.NewRepositories(db)
	services := service.NewServices(repos.Questions, repos.Decks, bot)

	tgHandler := handler.NewHandler(services.TgMessages, services.Cache)

	for update := range updates {
		if update.Message != nil {
			reply, err := tgHandler.HandleMessageAndReply(update)
			if err != nil {
				println(err.Error())
				reply = sendUnknownCommandAnswer(update)
			}
			if reply != nil {
				_, err := bot.Instance.Send(reply)
				if err != nil {
					println("ERROR WHILE SENDING MESSAGE!")
					println(err)
				}
			}
		}
	}
}

func sendUnknownCommandAnswer(update tgbotapi.Update) *tgbotapi.MessageConfig {
	println("UnknownCommand")
	ans := tgbotapi.NewMessage(update.Message.Chat.ID, "Не совсем понял команду, либо произошла ошибка(\r\nПопробуй заново /start")
	return &ans
}
