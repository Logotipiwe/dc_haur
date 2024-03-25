package repo

import (
	"dc_haur/src/internal/model"
	"errors"
	"github.com/jinzhu/gorm"
)

type DeckLikes struct {
	db *gorm.DB
}

func NewDeckLikes(db *gorm.DB) *DeckLikes {
	return &DeckLikes{db: db}
}

func (r DeckLikes) Save(l *model.DeckLike) error {
	err := r.db.Save(l).Error
	return err
}

func (r DeckLikes) GetByUserAndDeck(userID, deckID string) (*model.DeckLike, error) {
	if userID == "" {
		return nil, errors.New("userId is empty when getting deckLike")
	}
	if deckID == "" {
		return nil, errors.New("deckId is empty when getting deckLike")
	}
	var res model.DeckLike
	where := model.DeckLike{DeckID: deckID, UserID: userID}
	err := r.db.Find(&res, where).Error
	return &res, err
}

func (r DeckLikes) DeleteByUserAndDeckId(userId string, deckId string) error {
	where := model.DeckLike{UserID: userId, DeckID: deckId}
	err := r.db.Debug().Delete(where, where).Error
	return err
}

func (r DeckLikes) GetAllLikesByUser(userId string) ([]*model.DeckLike, error) {
	if userId == "" {
		return nil, errors.New("userId is empty when getting question likes")
	}
	var res []*model.DeckLike
	where := model.DeckLike{UserID: userId}
	err := r.db.Find(&res, where).Error
	return res, err
}
