package model

type UsedQuestion struct {
	QuestionID string `gorm:"primary_key"`
	ClientID   string `gorm:"primary_key"`
}

func (u UsedQuestion) TableName() string {
	return "used_questions"
}
