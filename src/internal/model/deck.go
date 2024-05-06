package model

type Deck struct {
	ID           string  `gorm:"column:id" json:"id,omitempty"`
	LanguageCode string  `gorm:"column:language_code" json:"languageCode"`
	Name         string  `gorm:"column:name" json:"name,omitempty"`
	Emoji        *string `gorm:"column:emoji" json:"emoji"`
	Description  string  `gorm:"column:description" json:"description,omitempty"`
	Labels       string  `gorm:"column:labels;serializer:semicolonSeparated" json:"labels"`
	ImageID      string  `gorm:"column:vector_image_id" json:"image_id"`
	Hidden       bool    `gorm:"column:hidden" json:"hidden"`
	Promo        string  `gorm:"column:promo" json:"promo"`
}

func (Deck) TableName() string {
	return "decks"
}
