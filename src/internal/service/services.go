package service

import (
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/repo"
	servicedomain "dc_haur/src/internal/service/domain"
)

type Services struct {
	Cache            *CacheService
	TgKeyboard       *TgKeyboardService
	TgMessages       *TgMessageService
	TgUpdatesHandler *Handler
	TgBotInteractor  domain.BotInteractor
	Decks            *DecksService
	Questions        *QuestionsService
	Repos            *repo.Repositories
	*servicedomain.DecksLikesService
	*servicedomain.QuestionLikesService
}

func NewServices(repos *repo.Repositories, bot domain.BotInteractor) *Services {
	cache := NewCacheService()
	tgKeyboard := NewTgKeyboardsService(repos)
	tgMessages := NewTgMessageService(*tgKeyboard, *cache, bot, repos)
	tgHandler := NewHandler(tgMessages, cache)
	deckLikesService := servicedomain.NewDeckLikesService(repos)
	questionLikesService := servicedomain.NewQuestionLikesService(repos)
	return &Services{
		Cache:                cache,
		TgKeyboard:           tgKeyboard,
		TgMessages:           tgMessages,
		TgUpdatesHandler:     tgHandler,
		TgBotInteractor:      bot,
		Decks:                NewDecksService(repos),
		Questions:            NewQuestionsService(repos),
		DecksLikesService:    deckLikesService,
		QuestionLikesService: questionLikesService,
		Repos:                repos,
	}
}
