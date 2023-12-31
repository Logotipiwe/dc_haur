package repos

import (
	"dc_haur/src/internal/domain"
	"errors"
)

type MockHistoryRepoWithError struct{}

func (m *MockHistoryRepoWithError) Truncate() error {
	return errors.New("planned error")
}

func NewMockHistoryRepoWithError() *MockHistoryRepoWithError {
	return &MockHistoryRepoWithError{}
}

func (m *MockHistoryRepoWithError) Insert(i int64, question *domain.Question) error {
	return errors.New("planned error")
}
