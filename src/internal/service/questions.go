package service

import (
	"dc_haur/src/internal/model"
	"dc_haur/src/internal/repo"
	config "github.com/logotipiwe/dc_go_config_lib"
)

type QuestionsService struct {
	repo  *repo.Questions
	Repos *repo.Repositories
}

func NewQuestionsService(repos *repo.Repositories) *QuestionsService {
	return &QuestionsService{repos.Questions, repos}
}

func (q QuestionsService) GetRandQuestion(levelID string, clientId string) (question *model.Question, isLast bool, err error) {
	dumbShuffle, errDumb := config.GetConfigBool("IS_DUMB_SHUFFLE")
	if errDumb == nil && dumbShuffle == true {
		question, err := q.repo.GetRandQuestionDumb(levelID)
		return question, false, err
	}

	question, err = q.repo.GetRandQuestionByLevel(levelID, clientId)
	if err != nil {
		return nil, false, err
	}
	err = q.Repos.History.Insert(clientId, question)
	if err != nil {
		return nil, false, err
	}
	err = q.Repos.UsedQuestions.Insert(clientId, question)
	if err != nil {
		return nil, false, err
	}
	allUsed, err := q.Repos.UsedQuestions.AreAllQuestionsUsed(levelID, clientId)
	if err != nil {
		return nil, false, err
	}
	if allUsed {
		err = q.Repos.UsedQuestions.Clear(levelID, clientId)
		if err != nil {
			return nil, false, err
		}
	}
	return question, allUsed, nil
}

func (q QuestionsService) GetLevelByNames(deckName string, levelNameWithEmoji string) (*model.Level, error) {
	level, err := q.repo.GetLevelByNames(deckName, levelNameWithEmoji)
	return level, err
}
