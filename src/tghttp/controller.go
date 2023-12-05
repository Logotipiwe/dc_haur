package tghttp

import (
	"database/sql"
	handler "dc_haur/src/internal"
	"dc_haur/src/internal/repo"
	"dc_haur/src/internal/service"
	"dc_haur/src/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/logotipiwe/dc_go_utils/src/config"
	"log"
)

func Start() {
	err, db := initializeApp()
	tgBot(db)
	println("Tg bot started!")
	println("Server up!")
	if err != nil {
		panic("Lol server fell")
	}
}

func tgBot(db *sql.DB) {
	bot, err := tgbotapi.NewBotAPI(config.GetConfig("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	repos := repo.NewRepositories(db)
	services := service.NewServices(repos.Questions, repos.Decks)

	tgHandler := handler.NewHandler(services.TgMessages, services.Cache)

	for update := range updates {
		if update.Message != nil {
			err, reply := tgHandler.HandleMessageAndReply(update)
			if err != nil {
				println(err.Error())
				reply = sendUnknownCommandAnswer(update)
			}
			if reply != nil {
				_, err := bot.Send(*reply)
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

func initializeApp() (error, *sql.DB) {
	config.LoadDcConfig()
	err, db := pkg.InitDb()
	if err != nil {
		panic(err)
	}
	return err, db
}
