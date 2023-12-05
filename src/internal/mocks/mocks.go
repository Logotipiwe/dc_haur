package mocks

import "dc_haur/src/internal/mocks/repos"

type Mocks struct {
	QuestionRepo        *repos.MockQuestionsRepo
	QuestionRepoWithErr *repos.MockQuestionsRepoWithError
	DeckRepo            *repos.MockDecksRepo
	DeckRepoWithErr     *repos.MockDecksRepoWithError
}

func NewMocks() Mocks {
	return Mocks{
		QuestionRepo:        repos.NewMockQuestionsRepo(),
		QuestionRepoWithErr: repos.NewMockQuestionsRepoWithError(),
		DeckRepo:            repos.NewMockDecksRepo(),
		DeckRepoWithErr:     repos.NewMockDecksRepoWithError(),
	}
}
