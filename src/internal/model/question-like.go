package model

type QuestionLike struct {
	ID         string `gorm:"column:id" json:"id,omitempty"`
	QuestionID string `gorm:"column:question_id" json:"questionId"`
	UserID     string `gorm:"column:user_id" json:"userId"`
}

func (d QuestionLike) TableName() string {
	return "question_likes"
}
