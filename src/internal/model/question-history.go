package model

type QuestionHistory struct {
	ID         string `gorm:"column:id"`
	LevelID    string `gorm:"column:level_id"`
	QuestionID string `gorm:"column:question_id"`
	ClientID   string `gorm:"column:client_id"`
}

func (QuestionHistory) TableName() string {
	return "questions_history"
}
