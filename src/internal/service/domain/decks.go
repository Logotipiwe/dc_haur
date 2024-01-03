package domain

import (
	"dc_haur/src/internal/repo"
)

type DecksService struct {
	repo.Decks
	Repos *repo.Repositories
}

func NewDecksService(repos *repo.Repositories) *DecksService {
	return &DecksService{repos.Decks, repos}
}
