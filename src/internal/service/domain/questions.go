package domain

import (
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/repo"
)

type QuestionsService struct {
	Repos *repo.Repositories
}

func NewQuestionsService(repos *repo.Repositories) *QuestionsService {
	return &QuestionsService{repos}
}

func (q QuestionsService) GetLevels(deckID string) ([]string, error) {
	levels, err := q.Repos.Questions.GetLevels(deckID)
	return levels, err
}

func (q QuestionsService) GetRandQuestion(deckID, levelName string) (*domain.Question, error) {
	return q.Repos.Questions.GetRandQuestion(deckID, levelName)
}

func (q QuestionsService) GetLevelsByName(deckName string) ([]string, error) {
	return q.Repos.Questions.GetLevelsByName(deckName)
}

func (q QuestionsService) GetRandQuestionByNames(deckName string, levelName string) (*domain.Question, error) {
	return q.Repos.Questions.GetRandQuestionByNames(deckName, levelName)
}
