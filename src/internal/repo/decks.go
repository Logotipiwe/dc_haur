package repo

import (
	"dc_haur/src/internal/domain"
	"github.com/jinzhu/gorm"
)

type Decks struct {
	db *gorm.DB
}

func NewDecksRepo(db *gorm.DB) *Decks {
	return &Decks{db: db}
}

func (r *Decks) GetDecks() ([]domain.Deck, error) {
	var decks []domain.Deck
	if err := r.db.Find(&decks).Error; err != nil {
		return nil, err
	}
	return decks, nil
}

func (r *Decks) GetDeckByName(name string) (*domain.Deck, error) {
	var deck domain.Deck
	if err := r.db.Where("name = ?", name).First(&deck).Error; err != nil {
		return nil, err
	}
	return &deck, nil
}

func (r *Decks) GetDeckByNameWithEmoji(name string) (*domain.Deck, error) {
	var deck domain.Deck
	if err := r.db.Where("(concat(coalesce(concat(emoji, ' '),''), name) = ?) OR name = ?", name, name).First(&deck).Error; err != nil {
		return nil, err
	}
	return &deck, nil
}
