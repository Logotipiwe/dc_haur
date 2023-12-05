package service

import (
	"dc_haur/src/internal/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tgkbTestRepoMocks = mocks.NewMocks()

func TestGetLevelsKeyboard(t *testing.T) {
	service := NewTgKeyboardsService(tgkbTestRepoMocks.QuestionRepo, tgkbTestRepoMocks.DeckRepo)
	err, keyboard := service.GetLevelsKeyboard("TestDeck")
	assert.Nil(t, err)
	assert.NotNil(t, keyboard)
	assert.True(t, len(keyboard.Keyboard) > 0)
}

func TestGetDecksKeyboard(t *testing.T) {
	service := NewTgKeyboardsService(tgkbTestRepoMocks.QuestionRepo, tgkbTestRepoMocks.DeckRepo)
	err, keyboard := service.GetDecksKeyboard()
	assert.Nil(t, err)
	assert.NotNil(t, keyboard)
	assert.True(t, len(keyboard.Keyboard) > 0)
}

func TestGetLevelsKeyboardWithError(t *testing.T) {
	service := NewTgKeyboardsService(tgkbTestRepoMocks.QuestionRepoWithErr, tgkbTestRepoMocks.DeckRepo)
	err, _ := service.GetLevelsKeyboard("TestDeck")
	assert.NotNil(t, err)
}

func TestGetDecksKeyboardWithError(t *testing.T) {
	service := NewTgKeyboardsService(tgkbTestRepoMocks.QuestionRepo, tgkbTestRepoMocks.DeckRepoWithErr)
	err, _ := service.GetDecksKeyboard()
	assert.NotNil(t, err)
}
