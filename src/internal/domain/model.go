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

type Bot struct {
	Instance *tgbotapi.BotAPI
}

func (bot Bot) SendToOwner(text string) error {
	ownerChatIdStr := utils.GetOwnerChatID()
	if ownerChatIdStr == "" {
		return errors.New("empty owner chat, unable to send message")
	}
	ownerChatID, err := strconv.ParseInt(ownerChatIdStr, 10, 64)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(ownerChatID, text)
	_, err = bot.Instance.Send(msg)
	return err
}
