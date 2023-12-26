package mocks

import (
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/mocks/repos"
	"dc_haur/src/internal/repo"
)

func NewRepoMocks() *repo.Repositories {
	return &repo.Repositories{
		Questions: repos.NewMockQuestionsRepo(),
		Decks:     repos.NewMockDecksRepo(),
		History:   repos.NewMockHistoryRepo(),
	}
}

func NewRepoMocksWithErrors() repo.Repositories {
	return repo.Repositories{
		Questions: repos.NewMockQuestionsRepoWithError(),
		Decks:     repos.NewMockDecksRepoWithError(),
	}
}

func NewBotInteractorMock() domain.BotInteractor {
	return repos.BotInteractorMock{}
}
