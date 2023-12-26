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
	Repos            *repo.Repositories
}

func NewServices(repos *repo.Repositories, bot domain.BotInteractor) *Services {
	cache := NewCacheService()
	tgKeyboard := NewTgKeyboardsService(repos)
	tgMessages := NewTgMessageService(*tgKeyboard, *cache, bot, repos)
	tgHandler := NewHandler(tgMessages, cache)
	return &Services{
		Cache:            cache,
		TgKeyboard:       tgKeyboard,
		TgMessages:       tgMessages,
		TgUpdatesHandler: tgHandler,
		TgBotInteractor:  bot,
		Repos:            repos,
	}
}
