package repo

import (
	"database/sql"
	"dc_haur/src/internal/domain"
)

type DecksRepo struct {
	db *sql.DB
}

func NewDecksRepo(db *sql.DB) *DecksRepo {
	return &DecksRepo{db: db}
}

func (r *DecksRepo) GetDecks() ([]domain.Deck, error) {
	res, err := r.db.Query(`SELECT id, name, description FROM decks`)
	if err != nil {
		return nil, err
	}
	var decks = make([]domain.Deck, 0)
	for res.Next() {
		deck := domain.Deck{}
		err := res.Scan(&deck.ID, &deck.Name, &deck.Description)
		if err != nil {
			return nil, err
		}
		decks = append(decks, deck)
	}
	return decks, nil
}
