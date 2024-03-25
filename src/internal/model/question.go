package model

type Question struct {
	ID             string  `gorm:"column:id" json:"id"`
	LevelID        string  `gorm:"column:level_id" json:"level_id"`
	Text           string  `gorm:"column:text" json:"text"`
	AdditionalText *string `gorm:"column:additional_text" json:"additional_text"`
}

func (Question) TableName() string {
	return "questions"
}
