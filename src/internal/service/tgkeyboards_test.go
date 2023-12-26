package service

import (
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/mocks"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tgkbTestRepoMocks = mocks.NewRepoMocks()

func TestGetLevelsKeyboard(t *testing.T) {
	levels := []string{"a", "b", "c", "d"}
	expected := tgbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tgbotapi.KeyboardButton{
			{
				tgbotapi.NewKeyboardButton("a"),
				tgbotapi.NewKeyboardButton("b"),
				tgbotapi.NewKeyboardButton("c"),
			},
			{
				tgbotapi.NewKeyboardButton("d"),
			},
		},
		ResizeKeyboard:        true,
		OneTimeKeyboard:       false,
		InputFieldPlaceholder: "",
		Selective:             false,
	}
	service := NewTgKeyboardsService(tgkbTestRepoMocks)
	keyboard := service.GetLevelsKeyboard(levels)
	assert.Equal(t, expected, keyboard)
}

func TestGetDecksKeyboard(t *testing.T) {
	decks := []domain.Deck{
		{
			ID:          "",
			Name:        "1",
			Description: "",
		},
		{
			ID:          "",
			Name:        "2",
			Description: "",
		},
		{
			ID:          "",
			Name:        "3",
			Description: "",
		},
		{
			ID:          "",
			Name:        "4",
			Description: "",
		},
	}
	expected := tgbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tgbotapi.KeyboardButton{
			{
				tgbotapi.NewKeyboardButton("1"),
			},
			{
				tgbotapi.NewKeyboardButton("2"),
			},
			{
				tgbotapi.NewKeyboardButton("3"),
			},
			{
				tgbotapi.NewKeyboardButton("4"),
			},
		},
		ResizeKeyboard:        true,
		OneTimeKeyboard:       false,
		InputFieldPlaceholder: "",
		Selective:             false,
	}
	service := NewTgKeyboardsService(tgkbTestRepoMocks)
	keyboard := service.GetDecksKeyboard(decks)
	assert.Equal(t, expected, keyboard)
}
