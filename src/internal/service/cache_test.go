package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestCacheService(t *testing.T) {
	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: 12345,
			},
		},
	}

	tests := []struct {
		name  string
		run   func(cs *CacheService)
		check func(cs *CacheService) bool
	}{
		{
			name: "AssignDeckToChat",
			run: func(cs *CacheService) {
				cs.AssignDeckToChat(update, "testDeck")
			},
			check: func(cs *CacheService) bool {
				found, deck := cs.GetCurrentChatDeckName(update)
				return found && deck == "testDeck"
			},
		},
		{
			name: "RemoveDeckFromChat",
			run: func(cs *CacheService) {
				cs.RemoveDeckFromChat(update)
			},
			check: func(cs *CacheService) bool {
				found, _ := cs.GetCurrentChatDeckName(update)
				return !found
			},
		},
		{
			name: "AssignFeedbackToChat",
			run: func(cs *CacheService) {
				cs.AssignFeedbackToChat(update)
			},
			check: func(cs *CacheService) bool {
				return cs.IsChatFeedbacks(update)
			},
		},
		{
			name: "RemoveFeedbackFromChat",
			run: func(cs *CacheService) {
				cs.RemoveFeedbackFromChat(update)
			},
			check: func(cs *CacheService) bool {
				return !cs.IsChatFeedbacks(update)
			},
		},
		{
			name: "AssignNewQuestionToChat",
			run: func(cs *CacheService) {
				cs.AssignNewQuestionToChat(update)
			},
			check: func(cs *CacheService) bool {
				return cs.IsChatNewQuestions(update)
			},
		},
		{
			name: "RemoveNewQuestionFromChat",
			run: func(cs *CacheService) {
				cs.RemoveNewQuestionFromChat(update)
			},
			check: func(cs *CacheService) bool {
				return !cs.IsChatNewQuestions(update)
			},
		},
		{
			name: "RemoveChatFromCaches",
			run: func(cs *CacheService) {
				cs.RemoveChatFromCaches(update)
			},
			check: func(cs *CacheService) bool {
				isFeedback := cs.IsChatFeedbacks(update)
				isNewQuestion := cs.IsChatNewQuestions(update)
				found, _ := cs.GetCurrentChatDeckName(update)
				return !found && !isNewQuestion && !isFeedback
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cs := NewCacheService()
			test.run(cs)
			if !test.check(cs) {
				t.Errorf("test %s failed", test.name)
			}
		})
	}
}
