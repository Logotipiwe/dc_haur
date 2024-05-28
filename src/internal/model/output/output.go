package output

import "dc_haur/src/internal/model"

type UserReactionsDTO struct {
	DeckLikes         []model.DeckLike         `json:"deckLikes"`
	QuestionReactions []model.QuestionReaction `json:"questionReactions"`
}

func NewUserReactions(deckLikes []model.DeckLike, questionReactions []model.QuestionReaction) UserReactionsDTO {
	return UserReactionsDTO{
		DeckLikes:         deckLikes,
		QuestionReactions: questionReactions,
	}
}
