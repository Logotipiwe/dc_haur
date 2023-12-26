package repo

import (
	"dc_haur/src/internal/domain"
	"github.com/jinzhu/gorm"
)

type DecksRepo struct {
	db *gorm.DB
}

func NewDecksRepo(db *gorm.DB) *DecksRepo {
	return &DecksRepo{db: db}
}

func (r *DecksRepo) GetDecks() ([]domain.Deck, error) {
	var decks []domain.Deck
	if err := r.db.Find(&decks).Error; err != nil {
		return nil, err
	}
	return decks, nil
}
