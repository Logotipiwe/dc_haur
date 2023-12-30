package service

import (
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/repo"
	domain2 "dc_haur/src/internal/service/domain"
)

type Services struct {
	Cache            *CacheService
	TgKeyboard       *TgKeyboardService
	TgMessages       *TgMessageService
	TgUpdatesHandler *Handler
	TgBotInteractor  domain.BotInteractor
	Decks            *domain2.DecksService
	Questions        *domain2.QuestionsService
	Repos            *repo.Repositories
}

func NewServices(repos *repo.Repositories, bot domain.BotInteractor) *Services {
	cache := NewCacheService()
	tgKeyboard := NewTgKeyboardsService(repos)
	tgMessages := NewTgMessageService(*tgKeyboard, *cache, bot, repos)
	tgHandler := NewHandler(tgMessages, cache)
	decksService := domain2.NewDecksService(repos)
	questionsService := domain2.NewQuestionsService(repos)
	return &Services{
		Cache:            cache,
		TgKeyboard:       tgKeyboard,
		TgMessages:       tgMessages,
		TgUpdatesHandler: tgHandler,
		TgBotInteractor:  bot,
		Decks:            decksService,
		Questions:        questionsService,
		Repos:            repos,
	}
}
