package service

import (
	"dc_haur/src/internal/repo"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func (s *TgMessageService) GetQuestionMessage(update Update, deckName string, levelName string) (error, *MessageConfig) {
	println("getQuestionMessage")
	err, question := s.questionsRepo.GetRandQuestion(deckName, levelName)
	if err != nil {
		return err, nil
	}
	msg := NewMessage(update.Message.Chat.ID, question.Text)
	return nil, &msg
}
