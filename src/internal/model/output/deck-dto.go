package output

import "dc_haur/src/internal/model"

type DeckDTO struct {
	model.Deck
	CardsCount int `json:"cardsCount"`
}
