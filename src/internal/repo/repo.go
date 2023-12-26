package repo

import (
	"database/sql"
	"dc_haur/src/internal/domain"
)

type Decks interface {
	GetDecks() ([]domain.Deck, error)
}

type Questions interface {
	GetLevels(deckName string) ([]string, error)
	GetRandQuestion(deckName string, levelName string) (*domain.Question, error)
}

type History interface {
	Insert(int64, *domain.Question) error
	Truncate() error
}

type Repositories struct {
	Decks
	Questions
	History
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Decks:     NewDecksRepo(db),
		Questions: NewQuestionsRepo(db),
		History:   NewQuestionsHistoryRepo(db),
	}
}
