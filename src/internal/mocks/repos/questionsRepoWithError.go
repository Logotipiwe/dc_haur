package repos

import (
	"dc_haur/src/internal/domain"
	"errors"
)

type MockQuestionsRepoWithError struct{}

func NewMockQuestionsRepoWithError() *MockQuestionsRepoWithError {
	return &MockQuestionsRepoWithError{}
}

func (m *MockQuestionsRepoWithError) GetRandQuestionByNames(_ string, levelName string) (*domain.Question, error) {
	return nil, errors.New("err")
}

func (m *MockQuestionsRepoWithError) GetLevelsByName(_ string) ([]string, error) {
	return nil, errors.New("err")
}
