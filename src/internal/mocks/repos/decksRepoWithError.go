package repos

import (
	"dc_haur/src/internal/domain"
	"errors"
)

type MockDecksRepoWithError struct{}

func NewMockDecksRepoWithError() *MockDecksRepoWithError {
	return &MockDecksRepoWithError{}
}

func (m *MockDecksRepoWithError) GetDecks() (error, []domain.Deck) {
	return errors.New("err"), nil
}
