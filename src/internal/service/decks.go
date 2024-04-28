package service

import (
	"dc_haur/src/internal/model/output"
	"dc_haur/src/internal/repo"
)

type DecksService struct {
	*repo.Decks
	Repos *repo.Repositories
}

func NewDecksService(repos *repo.Repositories) *DecksService {
	return &DecksService{repos.Decks, repos}
}

func (s DecksService) GetDecksWithCardsCounts() ([]output.DeckDTO, error) {
	//TODO N+1 problem in this method
	decks, err := s.GetDecks()
	if err != nil {
		return nil, err
	}
	res := make([]output.DeckDTO, 0)
	for _, deck := range decks {
		cardsCount, err := s.GetQuestionsCount(deck.ID)
		if err != nil {
			return nil, err
		}
		deckDTO := output.DeckDTO{
			Deck:       deck,
			CardsCount: cardsCount,
		}
		res = append(res, deckDTO)
	}
	return res, nil
}
