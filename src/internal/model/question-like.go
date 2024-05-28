package model

type ReactionType string

const (
	LIKE    ReactionType = "LIKE"
	DISLIKE ReactionType = "DISLIKE"
)

type QuestionReaction struct {
	ID           string       `gorm:"column:id" json:"id,omitempty"`
	QuestionID   string       `gorm:"column:question_id" json:"questionId"`
	UserID       string       `gorm:"column:user_id" json:"userId"`
	ReactionType ReactionType `gorm:"column:reaction_type" json:"reactionType"`
}

func (d QuestionReaction) TableName() string {
	return "question_likes"
}
