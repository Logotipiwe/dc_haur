package repos

import "dc_haur/src/internal/domain"

type MockQuestionsRepo struct{}

func NewMockQuestionsRepo() *MockQuestionsRepo {
	return &MockQuestionsRepo{}
}

func (m *MockQuestionsRepo) GetRandQuestion(deckName string, levelName string) (error, *domain.Question) {
	return nil, &domain.Question{
		ID:     "1",
		Level:  "2",
		DeckID: "3",
		Text:   "some rand text",
	}
}

func (m *MockQuestionsRepo) GetLevels(deckName string) (error, []string) {
	return nil, []string{"Level1", "Level2", "Level3"}
}
