package main

import (
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	utils "github.com/logotipiwe/dc_go_utils/src"
)

func GetLevelsKeyboard(deckName string) (error, ReplyKeyboardMarkup) {
	err, levels := GetLevels(deckName)
	if err != nil {
		return err, ReplyKeyboardMarkup{}
	}
	levelsChunked := ChunkStrings(levels, 3)
	keyboard := utils.Map(levelsChunked, func(chunk []string) []KeyboardButton {
		return utils.Map(chunk, func(level string) KeyboardButton {
			return NewKeyboardButton(level)
		})
	})
	return nil, ReplyKeyboardMarkup{
		Keyboard:              keyboard,
		ResizeKeyboard:        true,
		OneTimeKeyboard:       false,
		InputFieldPlaceholder: "",
		Selective:             false,
	}
}

func GetDecksKeyboard() (error, ReplyKeyboardMarkup) {
	err, decks := GetDecks()
	if err != nil {
		return err, ReplyKeyboardMarkup{}
	}
	keyboard := utils.Map(decks, func(deck Deck) []KeyboardButton {
		return []KeyboardButton{NewKeyboardButton(deck.Name)}
	})
	return nil, ReplyKeyboardMarkup{
		Keyboard:              keyboard,
		ResizeKeyboard:        true,
		OneTimeKeyboard:       false,
		InputFieldPlaceholder: "",
		Selective:             false,
	}
}
