package service

import (
	"dc_haur/src/internal/model/output"
	"dc_haur/src/internal/repo"
)

type LevelsService struct {
	repo *repo.Levels
}

func NewLevelsService(repo *repo.Levels) *LevelsService {
	return &LevelsService{repo}
}

func (s *LevelsService) EnrichWithCounts(dto output.LevelDto, clientID string) (output.LevelDto, error) {
	dto, err := s.EnrichWithCount(dto)
	if err != nil {
		return dto, err
	}
	return s.EnrichWithOpenedCount(dto, clientID)
}

func (s *LevelsService) EnrichWithCount(dto output.LevelDto) (output.LevelDto, error) {
	count, err := s.repo.GetQuestionsCount(dto.ID)
	if err != nil {
		return dto, err
	}
	if dto.Counts == nil {
		dto.Counts = &output.QuestionsCounts{}
	}
	dto.Counts.QuestionsCount = count
	return dto, nil
}

func (s *LevelsService) EnrichWithOpenedCount(dto output.LevelDto, clientID string) (output.LevelDto, error) {
	count, err := s.repo.GetOpenedQuestionsCount(dto.ID, clientID)
	if err != nil {
		return dto, err
	}
	if dto.Counts == nil {
		dto.Counts = &output.QuestionsCounts{}
	}
	dto.Counts.OpenedQuestionsCount = &count
	return dto, nil
}
