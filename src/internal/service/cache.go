package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type CacheService struct {
	chatIdToDeckName    map[string]string
	chatIdToFeedback    map[string]bool
	chatIdToNewQuestion map[string]bool
}

func NewCacheService() *CacheService {
	return &CacheService{
		chatIdToDeckName:    make(map[string]string),
		chatIdToFeedback:    make(map[string]bool),
		chatIdToNewQuestion: make(map[string]bool),
	}
}

func (s *CacheService) AssignDeckToChat(update tgbotapi.Update, deckName string) {
	s.chatIdToDeckName[strconv.FormatInt(update.Message.Chat.ID, 10)] = deckName
}

func (s *CacheService) RemoveDeckFromChat(update tgbotapi.Update) {
	delete(s.chatIdToDeckName, strconv.FormatInt(update.Message.Chat.ID, 10))
}

func (s *CacheService) GetCurrentChatDeckName(update tgbotapi.Update) (found bool, deckName string) {
	deckName, exists := s.chatIdToDeckName[strconv.FormatInt(update.Message.Chat.ID, 10)]
	return exists, deckName
}

func (s *CacheService) AssignFeedbackToChat(update tgbotapi.Update) {
	s.chatIdToFeedback[strconv.FormatInt(update.Message.Chat.ID, 10)] = true
}

func (s *CacheService) RemoveFeedbackFromChat(update tgbotapi.Update) {
	delete(s.chatIdToFeedback, strconv.FormatInt(update.Message.Chat.ID, 10))
}

func (s *CacheService) IsChatFeedbacks(update tgbotapi.Update) (found bool) {
	suggest, exists := s.chatIdToFeedback[strconv.FormatInt(update.Message.Chat.ID, 10)]
	return exists && suggest
}

func (s *CacheService) AssignNewQuestionToChat(update tgbotapi.Update) {
	s.chatIdToNewQuestion[strconv.FormatInt(update.Message.Chat.ID, 10)] = true
}

func (s *CacheService) RemoveNewQuestionFromChat(update tgbotapi.Update) {
	delete(s.chatIdToNewQuestion, strconv.FormatInt(update.Message.Chat.ID, 10))
}

func (s *CacheService) IsChatNewQuestions(update tgbotapi.Update) (found bool) {
	res, exists := s.chatIdToNewQuestion[strconv.FormatInt(update.Message.Chat.ID, 10)]
	return exists && res
}

func (s *CacheService) RemoveChatFromCaches(update tgbotapi.Update) {
	s.RemoveDeckFromChat(update)
	s.RemoveFeedbackFromChat(update)
	s.RemoveNewQuestionFromChat(update)
}
