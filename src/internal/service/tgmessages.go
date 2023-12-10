package service

import (
	"dc_haur/src/internal/repo"
	"dc_haur/src/pkg"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	config "github.com/logotipiwe/dc_go_config_lib"
	"log"
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

const WelcomeMessage = "Привет! Это игра \"How Are You Really?\" на знакомство и сближение! Каждая колода имеет несколько уровней вопросов. Выбирай колоду которая понравится и бери вопросы комфортного для тебя уровня, чтобы приятно провести время двоем или в компании! \r\n\r\n Выбери колоду, чтобы начать!"

func (s *TgMessageService) HandleStart(update Update) (*MessageConfig, error) {
	message := update.Message

	decks, err := s.decksRepo.GetDecks()
	if err != nil {
		return nil, err
	}

	msg := NewMessage(message.Chat.ID, WelcomeMessage)
	msg.ReplyMarkup = s.keyboards.GetDecksKeyboard(decks)
	s.cache.RemoveDeckFromChat(update)
	return &msg, nil
}

func (s *TgMessageService) GetLevelsMessage(update Update, deckName string) (*MessageConfig, error) {
	log.Println("GetLevelsMessage")

	levels, err := s.questionsRepo.GetLevels(deckName)
	if err != nil {
		return nil, err
	}

	markup := s.keyboards.GetLevelsKeyboard(levels)

	message := NewMessage(update.Message.Chat.ID, "Вот твои уровни")
	message.ReplyMarkup = markup
	s.cache.AssignDeckToChat(update, deckName)
	return &message, nil
}

func (s *TgMessageService) GetQuestionMessage(update Update, deckName string, levelName string) (Chattable, error) {
	log.Println("GetQuestionMessage")

	question, err := s.questionsRepo.GetRandQuestion(deckName, levelName)
	if err != nil {
		return nil, err
	}

	if imagesEnabled() {
		cardImage, err := CreateImageCard(question.Text)
		if err != nil {
			return nil, err
		}
		bytes, err := utils.EncodeImageToBytes(cardImage)
		if err != nil {
			return nil, err
		}
		return PhotoConfig{
			BaseFile: BaseFile{
				BaseChat: BaseChat{ChatID: update.Message.Chat.ID},
				File:     FileBytes{Name: uuid.New().String() + ".jpg", Bytes: bytes},
			},
		}, nil
	} else {
		return NewMessage(update.Message.Chat.ID, question.Text), nil
	}
}

func imagesEnabled() bool {
	imagesEnabled, err := strconv.ParseBool(config.GetConfigOr("ENABLE_IMAGES", "false"))
	if err != nil {
		imagesEnabled = false
	}
	return imagesEnabled
}
