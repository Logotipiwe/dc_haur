package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/logotipiwe/dc_go_utils/src/config"
	"log"
)

func main() {
	err := initializeApp()
	tgBot()
	println("Tg bot started!")
	println("Server up!")
	if err != nil {
		panic("Lol server fell")
	}
}

func tgBot() {
	bot, err := tgbotapi.NewBotAPI(config.GetConfig("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			err, reply := HandleMessageAndReply(update)
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

func initializeApp() error {
	config.LoadDcConfig()
	err := InitDb()
	if err != nil {
		panic(err)
	}
	return err
}
