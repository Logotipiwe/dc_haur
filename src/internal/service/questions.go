package service

import "dc_haur/src/internal/repo"

type QuestionsService struct {
	*repo.Questions
	Repos *repo.Repositories
}

func NewQuestionsService(repos *repo.Repositories) *QuestionsService {
	return &QuestionsService{repos.Questions, repos}
}
