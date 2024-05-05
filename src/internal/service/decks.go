package service

import (
	"dc_haur/src/internal/model"
	"dc_haur/src/internal/model/output"
	"dc_haur/src/internal/repo"
	utils "dc_haur/src/pkg"
)

type DecksService struct {
	*repo.Decks
	Repos *repo.Repositories
}

func NewDecksService(repos *repo.Repositories) *DecksService {
	return &DecksService{repos.Decks, repos}
}

func (s DecksService) ToDtos(decks []model.Deck) []output.DeckDTO {
	return utils.Map(decks, func(d model.Deck) output.DeckDTO {
		return output.DeckDTO{Deck: d}
	})
}

func (s DecksService) EnrichDecksWithCounts(decks []model.Deck, clientID string) ([]output.DeckDTO, error) {
	dtos, err := s.EnrichDecksWithCardsCounts(s.ToDtos(decks))
	if err != nil {
		return nil, err
	}
	dtos, err = s.EnrichDecksWithOpenedCardsCounts(dtos, clientID)
	return dtos, err
}

func (s DecksService) EnrichDecksWithCardsCounts(dtos []output.DeckDTO) ([]output.DeckDTO, error) {
	//TODO N+1 problem in this method
	res := make([]output.DeckDTO, 0)
	for _, dto := range dtos {
		cardsCount, err := s.GetQuestionsCount(dto.ID)
		if err != nil {
			return nil, err
		}
		dto.CardsCount = cardsCount
		res = append(res, dto)
	}
	return res, nil
}

func (s DecksService) EnrichDecksWithOpenedCardsCounts(dtos []output.DeckDTO, clientID string) ([]output.DeckDTO, error) {
	//TODO N+1 problem in this method
	res := make([]output.DeckDTO, 0)
	for _, dto := range dtos {
		openedCount, err := s.GetOpenedQuestionsCount(dto.ID, clientID)
		if err != nil {
			return nil, err
		}
		dto.OpenedCount = openedCount
		res = append(res, dto)
	}
	return res, nil
}

func (s DecksService) TryUnlockDeck(promo, clientId string) (*model.Deck, error) {
	found := s.Decks.GetHiddenDeckByPromo(promo, clientId)
	if found != nil {
		err := s.Decks.UnlockDeck(found.ID, clientId)
		return found, err
	} else {
		return nil, nil
	}
}
