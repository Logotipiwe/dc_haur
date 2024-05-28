package domain

import (
	"dc_haur/src/internal/model"
	"dc_haur/src/internal/repo"
	"github.com/google/uuid"
)

type QuestionReactionsService struct {
	Repos *repo.Repositories
}

func NewQuestionReactionsService(repos *repo.Repositories) *QuestionReactionsService {
	return &QuestionReactionsService{repos}
}

func (d *QuestionReactionsService) Like(userId, questionId string) error {
	updated, err := d.updateIfExists(userId, questionId, model.LIKE)
	if updated || err != nil {
		return err
	}
	_, err = d.createReaction(userId, questionId, model.LIKE)
	return err
}

// TODO test for changing like to dislike

// DislikeDEPRECATED TODO remove later
func (d *QuestionReactionsService) DislikeDEPRECATED(userId, questionId string) error {
	err := d.Repos.QuestionReactions.DeleteByUserAndQuestionId(userId, questionId)
	return err
}

func (d *QuestionReactionsService) RemoveReaction(userId string, questionId string) error {
	return d.Repos.QuestionReactions.DeleteByUserAndQuestionId(userId, questionId)
}

func (d *QuestionReactionsService) DislikeV2(userId string, questionId string) error {
	updated, err := d.updateIfExists(userId, questionId, model.DISLIKE)
	if updated || err != nil {
		return err
	}
	_, err = d.createReaction(userId, questionId, model.DISLIKE)
	return err
}

func (d *QuestionReactionsService) updateIfExists(userId string, questionId string, reactionType model.ReactionType) (updated bool, err error) {
	question, err := d.Repos.QuestionReactions.FindByUserAndQuestion(userId, questionId)
	if err != nil {
		return false, err
	}
	if question != nil {
		question.ReactionType = reactionType
		return true, d.Repos.QuestionReactions.Save(question)
	}
	return false, nil
}

func (d *QuestionReactionsService) createReaction(userId string, questionId string, reactionType model.ReactionType) (model.QuestionReaction, error) {
	dislike := model.QuestionReaction{
		ID:           uuid.NewString(),
		QuestionID:   questionId,
		UserID:       userId,
		ReactionType: reactionType,
	}
	return dislike, d.Repos.QuestionReactions.Save(&dislike)
}
