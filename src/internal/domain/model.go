package domain

import (
	utils "dc_haur/src/pkg"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type Question struct {
	ID             string  `gorm:"column:id" json:"id"`
	LevelID        string  `gorm:"column:level_id" json:"level_id"`
	Text           string  `gorm:"column:text" json:"text"`
	AdditionalText *string `gorm:"column:additional_text" json:"additional_text"`
}

func (Question) TableName() string {
	return "questions"
}

type Deck struct {
	ID          string  `gorm:"column:id" json:"id,omitempty"`
	Name        string  `gorm:"column:name" json:"name,omitempty"`
	Emoji       *string `gorm:"column:emoji" json:"emoji"`
	Description string  `gorm:"column:description" json:"description,omitempty"`
	Labels      string  `gorm:"column:labels;serializer:semicolonSeparated" json:"labels"`
	Image       string  `gorm:"column:image" json:"image"`
}

func (Deck) TableName() string {
	return "decks"
}

type Level struct {
	ID          string `gorm:"primary_key;type:varchar(255);not null"`
	DeckID      string `gorm:"type:varchar(255);not null"`
	LevelOrder  int
	Name        string  `gorm:"type:varchar(255);not null"`
	Emoji       *string `gorm:"column:emoji" json:"emoji"`
	ColorStart  string  `gorm:"type:varchar(255);not null"`
	ColorEnd    string  `gorm:"type:varchar(255);not null"`
	ColorButton string  `gorm:"type:varchar(255);not null"`
}

type VectorImage struct {
	ID      string `gorm:"primary_key;" json:"id"`
	Content string `json:"content"`
}

func (v VectorImage) TableName() string {
	return "vector_images"
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
	LevelID    string `gorm:"column:level_id"`
	QuestionID string `gorm:"column:question_id"`
	ChatID     string `gorm:"column:chat_id"`
}

func (QuestionHistory) TableName() string {
	return "questions_history"
}
