package model

type Level struct {
	ID          string `gorm:"primary_key;type:varchar(255);not null"`
	DeckID      string `gorm:"type:varchar(255);not null"`
	LevelOrder  int
	Name        string  `gorm:"type:varchar(255);not null"`
	Emoji       *string `gorm:"column:emoji" json:"emoji"`
	ColorStart  string  `gorm:"type:varchar(255);not null"`
	ColorEnd    string  `gorm:"type:varchar(255);not null"`
	ColorButton string  `gorm:"type:varchar(255);not null"`
}
