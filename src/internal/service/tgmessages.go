package service

import (
	"dc_haur/src/internal/repo"
	"dc_haur/src/pkg"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/logotipiwe/dc_go_utils/src/config"
	"strconv"
)

const DefaultDeckName = "😉 Для пары"

type TgMessageService struct {
	keyboards     TgKeyboardService
	cache         CacheService
	questionsRepo repo.Questions
	decksRepo     repo.Decks
}

func NewTgMessageService(tgKeyboardService TgKeyboardService, cache CacheService, questions repo.Questions,
	decks repo.Decks) *TgMessageService {
	return &TgMessageService{
		keyboards:     tgKeyboardService,
		cache:         cache,
		questionsRepo: questions,
		decksRepo:     decks,
	}
}

func (s *TgMessageService) HandleStart(update Update) (error, *MessageConfig) {
	message := update.Message

	err, decks := s.decksRepo.GetDecks()
	if err != nil {
		return err, nil
	}

	msg := NewMessage(message.Chat.ID, "Привет! Это игра \"How Are You Really?\" на знакомство и сближение! Каждая колода имеет несколько уровней вопросов. Выбирай колоду которая понравится и бери вопросы комфортного для тебя уровня, чтобы приятно провести время двоем или в компании! \r\n\r\n Выбери колоду, чтобы начать!")
	msg.ReplyMarkup = s.keyboards.GetDecksKeyboard(decks)
	s.cache.RemoveDeckFromChat(update)
	return nil, &msg
}

func (s *TgMessageService) GetLevelsMessage(update Update, deckName string) (error, *MessageConfig) {
	println("GetLevelsMessage")

	err, levels := s.questionsRepo.GetLevels(deckName)
	if err != nil {
		return err, nil
	}

	markup := s.keyboards.GetLevelsKeyboard(levels)

	message := NewMessage(update.Message.Chat.ID, "Вот твои уровни")
	message.ReplyMarkup = markup
	s.cache.AssignDeckToChat(update, deckName)
	return nil, &message
}

func (s *TgMessageService) GetQuestionMessage(update Update, deckName string, levelName string) (error, Chattable) {
	println("getQuestionMessage")
	err, question := s.questionsRepo.GetRandQuestion(deckName, levelName)
	if err != nil {
		return err, nil
	}

	if imagesEnabled() {
		err, cardImage := CreateImageCard(question.Text)
		if err != nil {
			return err, nil
		}
		bytes, err := pkg.EncodeImageToBytes(cardImage)
		if err != nil {
			return err, nil
		}
		return nil, PhotoConfig{
			BaseFile: BaseFile{
				BaseChat: BaseChat{ChatID: update.Message.Chat.ID},
				File:     FileBytes{Name: uuid.New().String() + ".jpg", Bytes: bytes},
			},
			Caption: question.Text,
		}
	} else {
		return nil, NewMessage(update.Message.Chat.ID, question.Text)
	}
}

func imagesEnabled() bool {
	imagesEnabled, err := strconv.ParseBool(config.GetConfig("ENABLE_IMAGES"))
	if err != nil {
		imagesEnabled = false
	}
	return imagesEnabled
}
