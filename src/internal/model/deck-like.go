package model

type DeckLike struct {
	ID     string `gorm:"column:id" json:"id,omitempty"`
	DeckID string `gorm:"column:deck_id" json:"deckId"`
	UserID string `gorm:"column:user_id" json:"userId"`
}

func (d DeckLike) TableName() string {
	return "deck_likes"
}
