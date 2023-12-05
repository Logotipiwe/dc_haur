package repos

import "dc_haur/src/internal/domain"

type MockDecksRepo struct{}

func NewMockDecksRepo() *MockDecksRepo {
	return &MockDecksRepo{}
}

func (m *MockDecksRepo) GetDecks() (error, []domain.Deck) {
	return nil, []domain.Deck{{Name: "Deck1"}, {Name: "Deck2"}}
}
