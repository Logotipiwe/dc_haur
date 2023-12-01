package main

import (
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func HandleMessageAndReply(update Update) *MessageConfig {
	message := update.Message

	log.Printf("[%s] %s", message.From.UserName, message.Text)
	if message.Text == "/start" {
		return handleStart(update)
	} else {
		err, question := getQuestionMessage(update)
		if err != nil {
			println(err)
			return sendUnknownCommandAnswer(update)
		}
		return question
	}
}

func handleStart(update Update) *MessageConfig {
	println("StartCommand")
	message := update.Message
	msg := NewMessage(message.Chat.ID, "Привет! Это игра \"How Are You Really?\" на знакомство и сближение. Состоит из карточек с вопросами разных уровней. Выбирай подходящий уровень, зачитывай вопрос и отвечай на него. Можете устроить обсуждение и послушать других участников.")
	var levels []string
	err, levels := GetLevels()
	if err != nil {
		return sendUnknownCommandAnswer(update)
	}
	msg.ReplyMarkup = GetLevelsKeyboard(levels)
	return &msg
}
func sendUnknownCommandAnswer(update Update) *MessageConfig {
	println("UnknownCommand")
	ans := NewMessage(update.Message.Chat.ID, "Не совсем понял команду, либо произошла ошибка(")
	return &ans
}
func getQuestionMessage(update Update) (error, *MessageConfig) {
	println("getQuestionMessage")
	level := update.Message.Text
	err, question := GetRandQuestionByLevel(level)
	if err != nil {
		return err, nil
	}
	ans := NewMessage(update.Message.Chat.ID, question.Text)
	err, levels := GetLevels()
	if err != nil {
		return err, nil
	}
	ans.ReplyMarkup = GetLevelsKeyboard(levels)
	return nil, &ans
}

//TODO сделать описание и название колоды для юзеров
//TODO ранжирование вопросов с функцией сброса???
