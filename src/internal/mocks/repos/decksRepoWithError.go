package repos

import (
	"dc_haur/src/internal/domain"
	"errors"
)

type MockDecksRepoWithError struct{}

func NewMockDecksRepoWithError() *MockDecksRepoWithError {
	return &MockDecksRepoWithError{}
}

func (m *MockDecksRepoWithError) GetDecks() ([]domain.Deck, error) {
	return nil, errors.New("err")
}
