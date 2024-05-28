package repo

import (
	"dc_haur/src/internal/model"
	"errors"
	"github.com/jinzhu/gorm"
)

type QuestionReactions struct {
	db *gorm.DB
}

func NewQuestionLikes(db *gorm.DB) *QuestionReactions {
	return &QuestionReactions{db: db}
}

func (r QuestionReactions) Save(l *model.QuestionReaction) error {
	err := r.db.Save(l).Error
	return err
}

func (r QuestionReactions) FindByUserAndQuestion(userID, questionID string) (*model.QuestionReaction, error) {
	if userID == "" {
		return nil, errors.New("userId is empty when getting questionLike")
	}
	if questionID == "" {
		return nil, errors.New("questionId is empty when getting questionLike")
	}
	var res model.QuestionReaction
	where := model.QuestionReaction{QuestionID: questionID, UserID: userID}
	err := r.db.Find(&res, where).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &res, err
}

func (r QuestionReactions) DeleteByUserAndQuestionId(userId string, questionId string) error {
	where := model.QuestionReaction{UserID: userId, QuestionID: questionId}
	err := r.db.Debug().Delete(where, where).Error
	return err
}

func (r QuestionReactions) GetAllLikesByUserDEPRECATED(userId string) ([]*model.QuestionReaction, error) {
	if userId == "" {
		return nil, errors.New("userId is empty when getting question likes")
	}
	var res []*model.QuestionReaction
	where := model.QuestionReaction{UserID: userId}
	err := r.db.Find(&res, where).Error
	return res, err
}

func (r QuestionReactions) GetAllReactionsByUser(userId string) ([]model.QuestionReaction, error) {
	if userId == "" {
		return nil, errors.New("userId is empty when getting question reactions")
	}
	var res []model.QuestionReaction
	where := model.QuestionReaction{UserID: userId}
	err := r.db.Find(&res, where).Error
	return res, err
}

func (r QuestionReactions) Truncate() error {
	return r.db.Exec(`TRUNCATE TABLE question_likes`).Error
}
