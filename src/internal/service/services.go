package service

import (
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/repo"
)

type Services struct {
	Cache            *CacheService
	TgKeyboard       *TgKeyboardService
	TgMessages       *TgMessageService
	TgUpdatesHandler *Handler
	TgBotInteractor  domain.BotInteractor
}

func NewServices(questions repo.Questions, decks repo.Decks, bot domain.BotInteractor) *Services {
	cache := NewCacheService()
	tgKeyboard := NewTgKeyboardsService(questions, decks)
	tgMessages := NewTgMessageService(*tgKeyboard, *cache, questions, decks, bot)
	tgHandler := NewHandler(tgMessages, cache)
	return &Services{
		Cache:            cache,
		TgKeyboard:       tgKeyboard,
		TgMessages:       tgMessages,
		TgUpdatesHandler: tgHandler,
		TgBotInteractor:  bot,
	}
}
