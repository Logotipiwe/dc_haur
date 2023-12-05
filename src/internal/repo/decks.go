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

func (r *DecksRepo) GetDecks() (error, []domain.Deck) {
	res, err := r.db.Query(`SELECT id, name, description FROM decks`)
	if err != nil {
		return err, nil
	}
	var decks = make([]domain.Deck, 0)

	for res.Next() {
		deck := domain.Deck{}
		err := res.Scan(&deck.ID, &deck.Name, &deck.Description)
		if err != nil {
			return err, nil
		}
		decks = append(decks, deck)
	}
	return nil, decks
}
