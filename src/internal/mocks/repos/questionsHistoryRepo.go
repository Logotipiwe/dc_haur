package repos

import "dc_haur/src/internal/domain"

type MockHistoryRepo struct{}

func NewMockHistoryRepo() *MockHistoryRepo {
	return &MockHistoryRepo{}
}

func (m *MockHistoryRepo) Insert(i int64, question *domain.Question) error {
	return nil
}
