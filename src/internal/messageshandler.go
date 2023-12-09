package internal

import (
	"dc_haur/src/internal/service"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Handler struct {
	messagesService *service.TgMessageService
	cache           *service.CacheService
}

func NewHandler(messageService *service.TgMessageService, cacheService *service.CacheService) *Handler {
	return &Handler{
		messagesService: messageService,
		cache:           cacheService,
	}
}

func (h *Handler) HandleMessageAndReply(update Update) (error, Chattable) {
	message := update.Message

	log.Printf("[%s] %s", message.From.UserName, message.Text)
	if message.Text == "/start" {
		println("StartCommand")
		return h.messagesService.HandleStart(update)
	} else if found, deckName := h.cache.GetCurrentChatDeckName(update); found {
		return h.messagesService.GetQuestionMessage(update, deckName, update.Message.Text)
	} else {
		return h.messagesService.GetLevelsMessage(update, update.Message.Text)
	}
}
