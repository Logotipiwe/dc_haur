package main

import (
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	utils "github.com/logotipiwe/dc_go_utils/src"
)

func GetLevelsKeyboard(levels []string) ReplyKeyboardMarkup {
	levelsChunked := ChunkStrings(levels, 3)
	keyboard := utils.Map(levelsChunked, func(levelsRow []string) []KeyboardButton {
		return utils.Map(levelsRow, func(level string) KeyboardButton {
			return NewKeyboardButton(level)
		})
	})
	return ReplyKeyboardMarkup{
		Keyboard:              keyboard,
		ResizeKeyboard:        true,
		OneTimeKeyboard:       false,
		InputFieldPlaceholder: "",
		Selective:             false,
	}
}
