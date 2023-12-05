package repos

import (
	"dc_haur/src/internal/domain"
	"errors"
)

type MockQuestionsRepoWithError struct{}

func NewMockQuestionsRepoWithError() *MockQuestionsRepoWithError {
	return &MockQuestionsRepoWithError{}
}

func (m *MockQuestionsRepoWithError) GetRandQuestion(deckName string, levelName string) (error, *domain.Question) {
	return errors.New("err"), nil
}

func (m *MockQuestionsRepoWithError) GetLevels(deckName string) (error, []string) {
	return errors.New("err"), nil
}
