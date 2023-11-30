package main

import . "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetLevelsKeyboard() ReplyKeyboardMarkup {
	return ReplyKeyboardMarkup{
		Keyboard: [][]KeyboardButton{
			{
				NewKeyboardButton(QuestionsLevels[EASY]),
				NewKeyboardButton(QuestionsLevels[MEDIUM]),
				NewKeyboardButton(QuestionsLevels[HARD]),
			},
		},
		ResizeKeyboard:        true,
		OneTimeKeyboard:       false,
		InputFieldPlaceholder: "",
		Selective:             false,
	}
}
