package service

import (
	"dc_haur/src/internal/mocks"
	_ "errors"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock structs for testing
//type mockTgKeyboardService struct{}
//type mockCacheService struct{}
//type mockQuestionsRepo struct{}

/*func (m *mockTgKeyboardService) GetLevelsKeyboard(deckName string) (error, ReplyKeyboardMarkup) {
	// Implement mock behavior here
	return nil, ReplyKeyboardMarkup{
		Keyboard: [][]KeyboardButton{
			{
				NewKeyboardButton("lol1"),
				NewKeyboardButton("lol2"),
				NewKeyboardButton("lol3"),
			},
			{
				NewKeyboardButton("lol1"),
			},
		},
	}
}

func (m *mockCacheService) RemoveDeckFromChat(update Update) {
	// Implement mock behavior here
}

func (m *mockCacheService) AssignDeckToChat(update Update, deckName string) {
	// Implement mock behavior here
}

func (m *mockQuestionsRepo) GetRandQuestion(deckName string, levelName string) (error, domain.Question) {
	// Implement mock behavior here
	return nil, domain.Question{Text: "Test question"}
}*/

var tgmsgTestRepoMocks = mocks.NewMocks()
var services = NewServices(tgkbTestRepoMocks.QuestionRepo, tgkbTestRepoMocks.DeckRepo)

func TestHandleStart(t *testing.T) {
	// Create a new instance of TgMessageService with mock dependencies
	service := NewTgMessageService(*services.TgKeyboard, *services.Cache, tgkbTestRepoMocks.QuestionRepo)

	// Call the method you want to test
	err, messageConfig := service.HandleStart(Update{Message: &Message{Chat: &Chat{ID: 123}}})

	// Assert that there is no error
	assert.Nil(t, err)

	// Add more assertions based on your expected behavior
	assert.NotNil(t, messageConfig)
	assert.NotNil(t, messageConfig.ReplyMarkup)
}

func TestGetLevelsMessage(t *testing.T) {
	// Create a new instance of TgMessageService with mock dependencies
	service := NewTgMessageService(*services.TgKeyboard, *services.Cache, tgkbTestRepoMocks.QuestionRepo)

	// Call the method you want to test
	err, messageConfig := service.GetLevelsMessage(Update{Message: &Message{Chat: &Chat{ID: 123}}}, DefaultDeckName)

	// Assert that there is no error
	assert.Nil(t, err)

	// Add more assertions based on your expected behavior
	assert.NotNil(t, messageConfig)
	assert.NotNil(t, messageConfig.ReplyMarkup)
}

func TestGetQuestionMessage(t *testing.T) {
	// Create a new instance of TgMessageService with mock dependencies
	service := NewTgMessageService(*services.TgKeyboard, *services.Cache, tgkbTestRepoMocks.QuestionRepo)

	// Call the method you want to test
	err, messageConfig := service.GetQuestionMessage(Update{Message: &Message{Chat: &Chat{ID: 123}}}, DefaultDeckName, "Level1")

	// Assert that there is no error
	assert.Nil(t, err)

	// Add more assertions based on your expected behavior
	assert.NotNil(t, messageConfig)
	assert.NotNil(t, messageConfig.Text)
}
