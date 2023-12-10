package repos

import "dc_haur/src/internal/domain"

type MockDecksRepo struct{}

func NewMockDecksRepo() *MockDecksRepo {
	return &MockDecksRepo{}
}

func (m *MockDecksRepo) GetDecks() ([]domain.Deck, error) {
	return []domain.Deck{{Name: "Deck1"}, {Name: "Deck2"}}, nil
}
