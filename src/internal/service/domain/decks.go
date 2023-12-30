package domain

import (
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/repo"
)

type DecksService struct {
	Repos *repo.Repositories
}

func NewDecksService(repos *repo.Repositories) *DecksService {
	return &DecksService{Repos: repos}
}

func (s DecksService) GetDecks() ([]domain.Deck, error) {
	decks, err := s.Repos.Decks.GetDecks()
	return decks, err
}
