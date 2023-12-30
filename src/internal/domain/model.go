package domain

import (
	utils "dc_haur/src/pkg"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type Question struct {
	ID     string `gorm:"column:id" json:"id,omitempty"`
	Level  string `gorm:"column:level" json:"level,omitempty"`
	DeckID string `gorm:"column:deck_id" json:"deckId,omitempty"`
	Text   string `gorm:"column:text" json:"text,omitempty"`
}

func (Question) TableName() string {
	return "questions"
}

type Deck struct {
	ID          string `gorm:"column:id" json:"id,omitempty"`
	Name        string `gorm:"column:name" json:"name,omitempty"`
	Description string `gorm:"column:description" json:"description,omitempty"`
}

func (Deck) TableName() string {
	return "decks"
}

type BotInteractor interface {
	SendToOwner(text string) error
	GetBot() *tgbotapi.BotAPI
}

type TgBotInteractor struct {
	instance *tgbotapi.BotAPI
}

func (bot TgBotInteractor) GetBot() *tgbotapi.BotAPI {
	return bot.instance
}

func NewBotInteractor(instance *tgbotapi.BotAPI) BotInteractor {
	return TgBotInteractor{instance: instance}
}

func (bot TgBotInteractor) SendToOwner(text string) error {
	ownerChatIdStr := utils.GetOwnerChatID()
	if ownerChatIdStr == "" {
		return errors.New("empty owner chat, unable to send message")
	}
	ownerChatID, err := strconv.ParseInt(ownerChatIdStr, 10, 64)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(ownerChatID, text)
	_, err = bot.instance.Send(msg)
	return err
}

type QuestionHistory struct {
	ID         string `gorm:"column:id"`
	DeckID     string `gorm:"column:deck_id"`
	LevelName  string `gorm:"column:level_name"`
	QuestionID string `gorm:"column:question_id"`
	ChatID     string `gorm:"column:chat_id"`
}

func (QuestionHistory) TableName() string {
	return "questions_history"
}
