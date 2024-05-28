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
	Levels           *LevelsService
	Repos            *repo.Repositories
	*servicedomain.DecksLikesService
	*servicedomain.QuestionReactionsService
}

func NewServices(repos *repo.Repositories, bot domain.BotInteractor) *Services {
	cache := NewCacheService()
	tgKeyboard := NewTgKeyboardsService(repos)
	qService := NewQuestionsService(repos)
	tgMessages := NewTgMessageService(*tgKeyboard, *cache, bot, repos, qService)
	tgHandler := NewHandler(tgMessages, cache)
	deckLikesService := servicedomain.NewDeckLikesService(repos)
	questionReactionsService := servicedomain.NewQuestionReactionsService(repos)
	levelsService := NewLevelsService(repos.Levels)
	return &Services{
		Cache:                    cache,
		TgKeyboard:               tgKeyboard,
		TgMessages:               tgMessages,
		TgUpdatesHandler:         tgHandler,
		TgBotInteractor:          bot,
		Decks:                    NewDecksService(repos),
		Questions:                qService,
		DecksLikesService:        deckLikesService,
		QuestionReactionsService: questionReactionsService,
		Repos:                    repos,
		Levels:                   levelsService,
	}
}
