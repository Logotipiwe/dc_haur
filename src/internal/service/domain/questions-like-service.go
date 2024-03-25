package domain

import (
	"dc_haur/src/internal/model"
	"dc_haur/src/internal/repo"
	"github.com/google/uuid"
)

type QuestionLikesService struct {
	Repos *repo.Repositories
}

func NewQuestionLikesService(repos *repo.Repositories) *QuestionLikesService {
	return &QuestionLikesService{repos}
}

func (d *QuestionLikesService) Like(userId, questionId string) error {
	like := model.QuestionLike{
		ID:         uuid.NewString(),
		QuestionID: questionId,
		UserID:     userId,
	}
	err := d.Repos.QuestionLikes.Save(&like)
	return err
}
func (d *QuestionLikesService) Dislike(userId, questionId string) error {
	err := d.Repos.QuestionLikes.DeleteByUserAndQuestionId(userId, questionId)
	return err
}
