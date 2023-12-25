package service

import (
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Handler struct {
	messagesService *TgMessageService
	cache           *CacheService
}

func NewHandler(messageService *TgMessageService, cacheService *CacheService) *Handler {
	return &Handler{
		messagesService: messageService,
		cache:           cacheService,
	}
}

func (h *Handler) HandleMessageAndReply(update Update) (Chattable, error) {
	message := update.Message

	log.Printf("[%s] %s", message.From.UserName, message.Text)
	if message.Text == "/start" {
		println("StartCommand")
		return h.messagesService.HandleStart(update)
	} else if is, sends := h.commandsFeedbackOrSendsFeedback(update); is {
		if !sends {
			return h.messagesService.AcceptFeedbackCommand(update)
		} else {
			return h.messagesService.AcceptFeedback(update)
		}
	} else if is, sends = h.commandsQuestionOrSendsQuestion(update); is {
		if !sends {
			return h.messagesService.AcceptNewQuestionCommand(update)
		} else {
			return h.messagesService.AcceptNewQuestion(update)
		}
	} else if found, deckName := h.cache.GetCurrentChatDeckName(update); found {
		return h.messagesService.GetQuestionMessage(update, deckName, update.Message.Text)
	} else {
		return h.messagesService.GetLevelsMessage(update, update.Message.Text)
	}
}

func (h *Handler) commandsQuestionOrSendsQuestion(update Update) (is, isCommand bool) {
	isCommand = update.Message.Text == "/question"
	isSends := h.cache.IsChatNewQuestions(update)
	return isCommand || isSends, isSends && !isCommand
}

// TODO if user send one command after another - command accepts as message
func (h *Handler) commandsFeedbackOrSendsFeedback(update Update) (is, isCommand bool) {
	isCommand = update.Message.Text == "/feedback"
	isSends := h.cache.IsChatFeedbacks(update)
	return isCommand || isSends, isSends && !isCommand
}
