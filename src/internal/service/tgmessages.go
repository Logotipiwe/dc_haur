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
}

func NewTgMessageService(tgKeyboardService TgKeyboardService, cache CacheService, questions repo.Questions) *TgMessageService {
	return &TgMessageService{
		keyboards:     tgKeyboardService,
		cache:         cache,
		questionsRepo: questions,
	}
}

func (s *TgMessageService) HandleStart(update Update) (error, *MessageConfig) {
	println("StartCommand")
	message := update.Message

	msg := NewMessage(message.Chat.ID, "Привет! Это игра \"How Are You Really?\" на знакомство и сближение! Каждая колода имеет несколько уровней вопросов. Выбирай колоду которая понравится и бери вопросы комфортного для тебя уровня, чтобы приятно провести время двоем или в компании! \r\n\r\n Выбери колоду, чтобы начать!")
	//uncomment when disable many decks
	//msg := NewMessage(message.Chat.ID, "Привет! Это игра \"How Are You Really?\" на знакомство и сближение. Состоит из карточек с вопросами разных уровней. Выбирай подходящий уровень, зачитывай вопрос и отвечай на него. Можете устроить обсуждение и послушать других участников.")

	err, keyboard := s.keyboards.GetDecksKeyboard() //uncomment when enable many decks
	//err, keyboard := s.keyboards.GetLevelsKeyboard(DefaultDeckName)
	if err != nil {
		return err, nil
	}
	msg.ReplyMarkup = keyboard

	s.cache.RemoveDeckFromChat(update)

	return nil, &msg
}

func (s *TgMessageService) GetLevelsMessage(update Update, deckName string) (error, *MessageConfig) {
	println("GetLevelsMessage")
	message := NewMessage(update.Message.Chat.ID, "Вот твои уровни")
	err, markup := s.keyboards.GetLevelsKeyboard(deckName)
	if err != nil {
		return err, nil
	}
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
