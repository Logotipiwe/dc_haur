package domain

import (
	"dc_haur/src/internal/model"
	"dc_haur/src/internal/repo"
	"github.com/google/uuid"
)

type DecksLikesService struct {
	Repos *repo.Repositories
}

func NewDeckLikesService(repos *repo.Repositories) *DecksLikesService {
	return &DecksLikesService{repos}
}

func (d *DecksLikesService) Like(userId, deckId string) error {
	like := model.DeckLike{
		ID:     uuid.NewString(),
		DeckID: deckId,
		UserID: userId,
	}
	err := d.Repos.DeckLikes.Save(&like)
	return err
}
func (d *DecksLikesService) Dislike(userId, deckId string) error {
	err := d.Repos.DeckLikes.DeleteByUserAndDeckId(userId, deckId)
	return err
}
