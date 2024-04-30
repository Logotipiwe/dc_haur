package service

import (
	"dc_haur/src/internal/model"
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

func (s DecksService) GetDecksByLanguageWithCounts(langCode string) ([]output.DeckDTO, error) {
	decks, err := s.GetDecksByLanguage(langCode)
	if err != nil {
		return nil, err
	}
	return s.EnrichDecksWithCounts(decks)
}

func (s DecksService) EnrichDecksWithCounts(decks []model.Deck) ([]output.DeckDTO, error) {
	//TODO N+1 problem in this method
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
