package model

type QuestionHistory struct {
	ID         string `gorm:"column:id"`
	LevelID    string `gorm:"column:level_id"`
	QuestionID string `gorm:"column:question_id"`
	ChatID     string `gorm:"column:chat_id"`
}

func (QuestionHistory) TableName() string {
	return "questions_history"
}
