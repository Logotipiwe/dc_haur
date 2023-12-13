package service

import (
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/mocks"
	_ "errors"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tgmsgTestRepoMocks = mocks.NewMocks()
var services = NewServices(tgkbTestRepoMocks.QuestionRepo, tgkbTestRepoMocks.DeckRepo, domain.Bot{})

func TestHandleStart(t *testing.T) {
	service := NewTgMessageService(*services.TgKeyboard, *services.Cache, tgmsgTestRepoMocks.QuestionRepo, tgmsgTestRepoMocks.DeckRepo, domain.Bot{})
	messageConfig, err := service.HandleStart(Update{Message: &Message{Chat: &Chat{ID: 123}}})
	assert.Nil(t, err)
	assert.NotNil(t, messageConfig)
	//assert.NotNil(t, messageConfig.ReplyMarkup) // TODO cover both messages with picture and text
}

func TestGetLevelsMessage(t *testing.T) {
	service := NewTgMessageService(*services.TgKeyboard, *services.Cache, tgkbTestRepoMocks.QuestionRepo, tgmsgTestRepoMocks.DeckRepo, domain.Bot{})
	messageConfig, err := service.GetLevelsMessage(Update{Message: &Message{Chat: &Chat{ID: 123}}}, DefaultDeckName)
	assert.Nil(t, err)
	assert.NotNil(t, messageConfig)
	//assert.NotNil(t, messageConfig.ReplyMarkup)
}

func TestGetQuestionMessage(t *testing.T) {
	service := NewTgMessageService(*services.TgKeyboard, *services.Cache, tgkbTestRepoMocks.QuestionRepo, tgmsgTestRepoMocks.DeckRepo, domain.Bot{})
	messageConfig, err := service.GetQuestionMessage(Update{Message: &Message{Chat: &Chat{ID: 123}}}, DefaultDeckName, "Level1")
	assert.Nil(t, err)
	assert.NotNil(t, messageConfig)
}
