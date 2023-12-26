package domain

import (
	utils "dc_haur/src/pkg"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type Question struct {
	ID     string `db:"id"`
	Level  string `db:"level"`
	DeckID string `db:"deck_id"`
	Text   string `db:"text"`
}

type Deck struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
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
	ChatID     int64  `gorm:"column:chat_id"`
}

func (QuestionHistory) TableName() string {
	return "questions_history"
}
