package mocks

import (
	"database/sql"
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/mocks/repos"
	"dc_haur/src/internal/repo"
)

type RepoMocks struct {
	QuestionRepo        repo.Questions
	QuestionRepoWithErr repo.Questions
	DeckRepo            repo.Decks
	DeckRepoWithErr     repo.Decks
}

func NewRepoMocks() RepoMocks {
	return RepoMocks{
		QuestionRepo:        repos.NewMockQuestionsRepo(),
		QuestionRepoWithErr: repos.NewMockQuestionsRepoWithError(),
		DeckRepo:            repos.NewMockDecksRepo(),
		DeckRepoWithErr:     repos.NewMockDecksRepoWithError(),
	}
}

func NewRepoMocksWithDb(db *sql.DB) RepoMocks {
	return RepoMocks{
		QuestionRepo:        repo.NewQuestionsRepo(db),
		QuestionRepoWithErr: repo.NewQuestionsRepo(db),
		DeckRepo:            repo.NewDecksRepo(db),
		DeckRepoWithErr:     repo.NewDecksRepo(db),
	}
}

func NewBotInteractorMock() domain.BotInteractor {
	return repos.BotInteractorMock{}
}
