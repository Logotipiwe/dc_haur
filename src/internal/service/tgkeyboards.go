package service

import (
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/repo"
	"dc_haur/src/pkg"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgKeyboardService struct {
	questionsRepo repo.Questions
	decksRepo     repo.Decks
}

func NewTgKeyboardsService(repos *repo.Repositories) *TgKeyboardService {
	return &TgKeyboardService{
		questionsRepo: repos.Questions,
		decksRepo:     repos.Decks,
	}

}

func (s *TgKeyboardService) GetLevelsKeyboard(levels []string) ReplyKeyboardMarkup {
	levelsChunked := utils.ChunkStrings(levels, 3)
	keyboard := utils.Map(levelsChunked, func(chunk []string) []KeyboardButton {
		return utils.Map(chunk, func(level string) KeyboardButton {
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

func (s *TgKeyboardService) GetDecksKeyboard(decks []domain.Deck) ReplyKeyboardMarkup {
	keyboard := utils.Map(decks, func(deck domain.Deck) []KeyboardButton {
		return []KeyboardButton{NewKeyboardButton(deck.Name)}
	})
	return ReplyKeyboardMarkup{
		Keyboard:              keyboard,
		ResizeKeyboard:        true,
		OneTimeKeyboard:       false,
		InputFieldPlaceholder: "",
		Selective:             false,
	}
}
