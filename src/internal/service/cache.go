package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CacheService struct {
	chatIdToDeckName map[string]string
}

func NewCacheService() *CacheService {
	return &CacheService{
		chatIdToDeckName: make(map[string]string),
	}
}

func (s *CacheService) AssignDeckToChat(update tgbotapi.Update, deckName string) {
	//s.chatIdToDeckName[strconv.FormatInt(update.Message.Chat.ID, 10)] = deckName
}

func (s *CacheService) RemoveDeckFromChat(update tgbotapi.Update) {
	//delete(s.chatIdToDeckName, strconv.FormatInt(update.Message.Chat.ID, 10))
}

func (s *CacheService) GetCurrentChatDeckName(update tgbotapi.Update) (found bool, deckName string) {
	//deckName, exists := s.chatIdToDeckName[strconv.FormatInt(update.Message.Chat.ID, 10)]
	//return exists, deckName
	return true, DefaultDeckName
}
