package repos

import "dc_haur/src/internal/domain"

type MockQuestionsRepo struct{}

func NewMockQuestionsRepo() *MockQuestionsRepo {
	return &MockQuestionsRepo{}
}

func (m *MockQuestionsRepo) GetRandQuestionByNames(_ string, _ string) (*domain.Question, error) {
	return &domain.Question{
		ID:     "1",
		Level:  "2",
		DeckID: "3",
		Text:   "some rand text",
	}, nil
}

func (m *MockQuestionsRepo) GetLevelsByName(_ string) ([]string, error) {
	return []string{"Level1", "Level2", "Level3"}, nil
}
