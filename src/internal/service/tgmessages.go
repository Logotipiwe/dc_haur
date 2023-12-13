package service

import (
	"dc_haur/src/internal/domain"
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
	bot           domain.Bot
}

func NewTgMessageService(tgKeyboardService TgKeyboardService, cache CacheService, questions repo.Questions, decks repo.Decks, bot domain.Bot) *TgMessageService {
	return &TgMessageService{
		keyboards:     tgKeyboardService,
		cache:         cache,
		questionsRepo: questions,
		decksRepo:     decks,
		bot:           bot,
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
	s.cache.RemoveChatFromCaches(update)
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

func (s *TgMessageService) AcceptFeedbackCommand(update Update) (*MessageConfig, error) {
	msg := NewMessage(update.Message.Chat.ID, "Отправьте свой отзыв одним сообщением. Мы получим его и учтём в будущем ❤️")
	msg.ReplyMarkup = ReplyKeyboardRemove{RemoveKeyboard: true}
	s.cache.AssignFeedbackToChat(update)
	return &msg, nil
}

func (s *TgMessageService) AcceptFeedback(update Update) (*MessageConfig, error) {
	msg := NewMessage(update.Message.Chat.ID, "Спасибо за отзыв! Мы его учтем ❤️. Отправьте /start, чтобы играть дальше.")
	userLink := "@" + update.Message.From.UserName
	err := s.bot.SendToOwner("Фидбек от " + userLink + ".\r\n" + update.Message.Text)
	if err != nil {
		return nil, err
	}
	s.cache.RemoveFeedbackFromChat(update)
	return &msg, nil
}

func (s *TgMessageService) AcceptNewQuestionCommand(update Update) (*MessageConfig, error) {
	msg := NewMessage(update.Message.Chat.ID, "Отправьте свой вопрос одним сообщением. Мы получим его и добавим в колоду вопросов ❤️")
	msg.ReplyMarkup = ReplyKeyboardRemove{RemoveKeyboard: true}
	s.cache.AssignNewQuestionToChat(update)
	return &msg, nil
}

func (s *TgMessageService) AcceptNewQuestion(update Update) (*MessageConfig, error) {
	msg := NewMessage(update.Message.Chat.ID, "Спасибо за вопрос! Мы его добавим в колоду вопросов ❤️. Отправьте /start, чтобы играть дальше.")
	userLink := "@" + update.Message.From.UserName
	err := s.bot.SendToOwner("Предложенный вопрос от " + userLink + ".\r\n" + update.Message.Text)
	if err != nil {
		return nil, err
	}
	s.cache.RemoveNewQuestionFromChat(update)
	return &msg, nil
}

func imagesEnabled() bool {
	imagesEnabled, err := strconv.ParseBool(config.GetConfigOr("ENABLE_IMAGES", "false"))
	if err != nil {
		imagesEnabled = false
	}
	return imagesEnabled
}
