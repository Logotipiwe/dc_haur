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
const GotLevelsMessage = "Вот твои уровни"

type TgMessageService struct {
	keyboards TgKeyboardService
	cache     CacheService
	repos     *repo.Repositories
	bot       domain.BotInteractor
}

const (
	AcceptFeedbackText    = "Спасибо за отзыв! Мы его учтем ❤️. Отправьте /start, чтобы играть дальше."
	AcceptNewQuestionText = "Спасибо за вопрос! Мы его добавим в колоду вопросов ❤️. Отправьте /start, чтобы играть дальше."
	AssignNewQuestionText = "Отправьте свой вопрос одним сообщением. Мы получим его и добавим в колоду вопросов ❤️"
	AssignFeedbackText    = "Отправьте свой отзыв одним сообщением. Мы получим его и учтём в будущем ❤️"
	WelcomeMessage        = "Привет! Это игра \"How Are You Really?\" на знакомство и сближение! Каждая колода имеет несколько уровней вопросов. Выбирай колоду которая понравится и бери вопросы комфортного для тебя уровня, чтобы приятно провести время двоем или в компании! \r\n\r\n Выбери колоду, чтобы начать!"
)

func NewTgMessageService(tgKeyboardService TgKeyboardService, cache CacheService,
	bot domain.BotInteractor, repos *repo.Repositories) *TgMessageService {
	return &TgMessageService{
		keyboards: tgKeyboardService,
		cache:     cache,
		bot:       bot,
		repos:     repos,
	}
}

func (s *TgMessageService) HandleStart(update Update) (*MessageConfig, error) {
	message := update.Message

	decks, err := s.repos.Decks.GetDecks()
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

	levels, err := s.repos.Questions.GetLevelsByName(deckName)
	if err != nil {
		return nil, err
	}

	markup := s.keyboards.GetLevelsKeyboard(levels)

	message := NewMessage(update.Message.Chat.ID, GotLevelsMessage)
	message.ReplyMarkup = markup
	s.cache.AssignDeckToChat(update, deckName)
	return &message, nil
}

func (s *TgMessageService) GetQuestionMessage(update Update, deckName string, levelName string) (Chattable, error) {
	log.Println("GetQuestionMessage")

	chatID := update.Message.Chat.ID
	question, err := s.repos.Questions.GetRandQuestionByNames(deckName, levelName)
	if err != nil {
		return nil, err
	}

	var chattable Chattable
	if imagesEnabled() {
		cardImage, err := CreateImageCard(question.Text)
		if err != nil {
			return nil, err
		}
		bytes, err := utils.EncodeImageToBytes(cardImage)
		if err != nil {
			return nil, err
		}
		chattable = PhotoConfig{
			BaseFile: BaseFile{
				BaseChat: BaseChat{ChatID: chatID},
				File:     FileBytes{Name: uuid.New().String() + ".jpg", Bytes: bytes},
			},
		}
	} else {
		chattable = NewMessage(update.Message.Chat.ID, question.Text)
	}
	err = s.repos.History.Insert(chatID, question)
	if err != nil {
		return nil, err
	}
	return chattable, nil
}

func (s *TgMessageService) AcceptFeedbackCommand(update Update) (*MessageConfig, error) {
	msg := NewMessage(update.Message.Chat.ID, AssignFeedbackText)
	msg.ReplyMarkup = ReplyKeyboardRemove{RemoveKeyboard: true}
	s.cache.AssignFeedbackToChat(update)
	return &msg, nil
}

func (s *TgMessageService) AcceptFeedback(update Update) (*MessageConfig, error) {
	msg := NewMessage(update.Message.Chat.ID, AcceptFeedbackText)
	userLink := "@" + update.Message.From.UserName
	err := s.bot.SendToOwner("Фидбек от " + userLink + ".\r\n" + update.Message.Text)
	if err != nil {
		return nil, err
	}
	s.cache.RemoveFeedbackFromChat(update)
	return &msg, nil
}

func (s *TgMessageService) AcceptNewQuestionCommand(update Update) (*MessageConfig, error) {
	msg := NewMessage(update.Message.Chat.ID, AssignNewQuestionText)
	msg.ReplyMarkup = ReplyKeyboardRemove{RemoveKeyboard: true}
	s.cache.AssignNewQuestionToChat(update)
	return &msg, nil
}

func (s *TgMessageService) AcceptNewQuestion(update Update) (*MessageConfig, error) {
	msg := NewMessage(update.Message.Chat.ID, AcceptNewQuestionText)
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
