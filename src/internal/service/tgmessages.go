package service

import (
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/model"
	"dc_haur/src/internal/repo"
	"dc_haur/src/pkg"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	config "github.com/logotipiwe/dc_go_config_lib"
	"strconv"
)

const DefaultDeckName = "😉 Для пары"
const GotLevelsMessage = "Ниже - список уровней. Чтобы получить вопрос - жми на нужный уровень :)"

type TgMessageService struct {
	keyboards        TgKeyboardService
	cache            CacheService
	repos            *repo.Repositories
	bot              domain.BotInteractor
	questionsService *QuestionsService
}

const (
	AcceptFeedbackText    = "Спасибо за отзыв! Мы его учтем ❤️. Отправьте /start, чтобы играть дальше."
	AcceptNewQuestionText = "Спасибо за вопрос! Мы его добавим в колоду вопросов ❤️. Отправьте /start, чтобы играть дальше."
	AssignNewQuestionText = "Отправьте свой вопрос одним сообщением. Мы получим его и добавим в колоду вопросов ❤️"
	AssignFeedbackText    = "Отправьте свой отзыв одним сообщением. Мы получим его и учтём в будущем ❤️"
	WelcomeMessage        = `Привет! Это разговорная игра "How Are You Really?" с вопросами на знакомство и сближение людей! 

Этот бот - аналог физической колоды карт, где каждый из игроков сидящих в кругу по очереди берёт карту и отвечает на вопрос. Только здесь игрокам нужно по очереди передавать смартфон.

Играть можно не только на очной встрече, но и созвонившись онлайн - тогда либо вопросы зачитывает ведущий, либо каждый берёт вопрос на своем устройстве самостоятельно.

Вопросы разбиты на тематические колоды, а также на уровни глубины.
В течение игры рекомендуется выбирать уровни по нарастанию.

 Выбери колоду, чтобы начать!`
)

func NewTgMessageService(tgKeyboardService TgKeyboardService, cache CacheService,
	bot domain.BotInteractor, repos *repo.Repositories, questionsService *QuestionsService) *TgMessageService {
	return &TgMessageService{
		keyboards:        tgKeyboardService,
		cache:            cache,
		bot:              bot,
		repos:            repos,
		questionsService: questionsService,
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

func (s *TgMessageService) GetLevelsMessage(update Update, deckNameWithEmoji string) (*MessageConfig, error) {
	deck, err := s.repos.Decks.GetDeckByNameWithEmoji(deckNameWithEmoji)
	if err != nil {
		return nil, err
	}
	levels, err := s.repos.Levels.GetLevelsByDeckId(deck.ID)
	if err != nil {
		return nil, err
	}

	levelsNames := pkg.Map(levels, func(l *model.Level) string {
		if l.Emoji != nil {
			return *l.Emoji + " " + l.Name
		}
		return l.Name
	})

	markup := s.keyboards.GetLevelsKeyboard(levelsNames)

	message := NewMessage(update.Message.Chat.ID, deck.Description+"\n\n"+GotLevelsMessage)
	message.ReplyMarkup = markup
	s.cache.AssignDeckToChat(update, deck.Name)
	return &message, nil
}

func (s *TgMessageService) GetQuestionMessage(update Update, deckName string, levelNameWithEmoji string) (Chattable, error) {
	chatID := update.Message.Chat.ID
	level, err := s.questionsService.GetLevelByNames(deckName, levelNameWithEmoji)
	if err != nil {
		return nil, err
	}
	question, _, err := s.questionsService.GetRandQuestion(level.ID, strconv.FormatInt(chatID, 10))
	if err != nil {
		return nil, err
	}

	if imagesEnabled() {
		return createPhotoConfig(question, level, chatID)
	} else {
		return NewMessage(update.Message.Chat.ID, question.Text), nil
	}
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

func createPhotoConfig(question *model.Question, level *model.Level, chatID int64) (*PhotoConfig, error) {
	cardImage, err := CreateImageCardFromQuestion(question, level.ColorStart, level.ColorEnd)
	if err != nil {
		return nil, err
	}
	bytes, err := pkg.EncodeImageToBytes(cardImage)
	if err != nil {
		return nil, err
	}
	return &PhotoConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: chatID},
			File:     FileBytes{Name: uuid.New().String() + ".jpg", Bytes: bytes},
		},
	}, nil
}
