package repos

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type BotInteractorMock struct {
}

func (b BotInteractorMock) GetBot() *tgbotapi.BotAPI {
	panic("implement me")
}

func (b BotInteractorMock) SendToOwner(text string) error {
	log.Println("invoked mock SendToOwner with text: " + text)
	return nil
}
