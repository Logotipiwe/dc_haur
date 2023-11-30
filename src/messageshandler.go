package main

import (
	"database/sql"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func HandleMessageAndReply(update Update) *MessageConfig {
	message := update.Message

	log.Printf("[%s] %s", message.From.UserName, message.Text)
	if message.Text == "/start" {
		return handleStart(update)
	} else if ExistsInArr(message.Text, GetValues(QuestionsLevels)) {
		return sendQuestion(update, QuestionsLevelsReverted[message.Text])
	} else {
		return sendUnknownCommandAnswer(update)
	}
}

func handleStart(update Update) *MessageConfig {
	println("StartCommand")
	message := update.Message
	msg := NewMessage(message.Chat.ID, "Привет! Это игра \"How Are You Really?\" на знакомство и сближение. Состоит из карточек с вопросами разных уровней. Выбирай подходящий уровень, зачитывай вопрос и отвечай на него. Можете устроить обсуждение и послушать других участников.")
	msg.ReplyMarkup = GetLevelsKeyboard()
	return &msg
}
func sendUnknownCommandAnswer(update Update) *MessageConfig {
	println("UnknownCommand")
	ans := NewMessage(update.Message.Chat.ID, "Не совсем понял команду, попробуй другую")
	return &ans
}
func sendQuestion(update Update, level string) *MessageConfig {
	println("sendQuestion")
	err, question := GetRandQuestionByLevel(level)
	if err != nil {
		var errMsg string
		if err == sql.ErrNoRows {
			errMsg = "В этой колоде не оказалось вопросов такого уровня. Попробуй другой"
		} else {
			errMsg = "У меня чето сломалось, попробуй ещё раз"
		}
		ans := NewMessage(update.Message.Chat.ID, errMsg)
		return &ans
	}
	ans := NewMessage(update.Message.Chat.ID, question.Text)
	return &ans
}
