package repo

import (
	"dc_haur/src/internal/model"
	"errors"
	"github.com/jinzhu/gorm"
)

type QuestionLikes struct {
	db *gorm.DB
}

func NewQuestionLikes(db *gorm.DB) *QuestionLikes {
	return &QuestionLikes{db: db}
}

func (r QuestionLikes) Save(l *model.QuestionLike) error {
	err := r.db.Save(l).Error
	return err
}

func (r QuestionLikes) GetByUserAndQuestion(userID, questionID string) (*model.QuestionLike, error) {
	if userID == "" {
		return nil, errors.New("userId is empty when getting questionLike")
	}
	if questionID == "" {
		return nil, errors.New("questionId is empty when getting questionLike")
	}
	var res model.QuestionLike
	where := model.QuestionLike{QuestionID: questionID, UserID: userID}
	err := r.db.Find(&res, where).Error
	return &res, err
}

func (r QuestionLikes) DeleteByUserAndQuestionId(userId string, questionId string) error {
	where := model.QuestionLike{UserID: userId, QuestionID: questionId}
	err := r.db.Debug().Delete(where, where).Error
	return err
}

func (r QuestionLikes) GetAllLikesByUser(userId string) ([]*model.QuestionLike, error) {
	if userId == "" {
		return nil, errors.New("userId is empty when getting question likes")
	}
	var res []*model.QuestionLike
	where := model.QuestionLike{UserID: userId}
	err := r.db.Find(&res, where).Error
	return res, err
}
