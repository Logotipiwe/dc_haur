package service

import (
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/repo"
)

type Services struct {
	Cache      *CacheService
	TgKeyboard *TgKeyboardService
	TgMessages *TgMessageService
}

func NewServices(questions repo.Questions, decks repo.Decks, bot domain.Bot) *Services {
	cache := NewCacheService()
	tgKeyboard := NewTgKeyboardsService(questions, decks)
	tgMessages := NewTgMessageService(*tgKeyboard, *cache, questions, decks, bot)
	return &Services{
		Cache:      cache,
		TgKeyboard: tgKeyboard,
		TgMessages: tgMessages,
	}
}
