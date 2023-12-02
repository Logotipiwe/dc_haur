package main

import (
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

var chatIdToDeckName = make(map[string]string)

func HandleMessageAndReply(update Update) (error, *MessageConfig) {
	message := update.Message

	log.Printf("[%s] %s", message.From.UserName, message.Text)
	if message.Text == "/start" {
		return handleStart(update)
	} else if found, deckName := getCurrentChatDeckName(update); found {
		return getQuestionMessage(update, deckName, update.Message.Text)
	} else {
		return getLevelsMessage(update, update.Message.Text)
	}
}

func handleStart(update Update) (error, *MessageConfig) {
	println("StartCommand")
	message := update.Message
	msg := NewMessage(message.Chat.ID, "Привет! Это игра \"How Are You Really?\" на знакомство и сближение! Каждая колода имеет несколько уровней вопросов. Выбирай колоду которая понравится и бери вопросы комфортного для тебя уровня, чтобы приятно провести время двоем или в компании! \r\n\r\n Выбери колоду, чтобы начать!")
	err, keyboard := GetDecksKeyboard()
	if err != nil {
		return err, nil
	}
	msg.ReplyMarkup = keyboard

	removeDeckFromChat(update)

	return nil, &msg
}

func getLevelsMessage(update Update, deckName string) (error, *MessageConfig) {
	println("GetLevelsMessage")
	message := NewMessage(update.Message.Chat.ID, "Вот твои уровни")
	err, markup := GetLevelsKeyboard(deckName)
	if err != nil {
		return err, nil
	}
	message.ReplyMarkup = markup
	assignDeckToChat(update, deckName)
	return nil, &message
}

func assignDeckToChat(update Update, deckName string) {
	chatIdToDeckName[strconv.FormatInt(update.Message.Chat.ID, 10)] = deckName
}

func removeDeckFromChat(update Update) {
	delete(chatIdToDeckName, strconv.FormatInt(update.Message.Chat.ID, 10))
}

func getCurrentChatDeckName(update Update) (found bool, deckName string) {
	deckName, exists := chatIdToDeckName[strconv.FormatInt(update.Message.Chat.ID, 10)]
	return exists, deckName
}

func getQuestionMessage(update Update, deckName string, levelName string) (error, *MessageConfig) {
	println("getQuestionMessage")
	err, question := GetRandQuestion(deckName, levelName)
	if err != nil {
		return err, nil
	}
	msg := NewMessage(update.Message.Chat.ID, question.Text)
	return nil, &msg
}
