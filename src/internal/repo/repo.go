package repo

import (
	"database/sql"
	"dc_haur/src/internal/domain"
)

type Decks interface {
	GetDecks() (error, []domain.Deck)
}

type Questions interface {
	GetLevels(deckName string) (error, []string)
	GetRandQuestion(deckName string, levelName string) (error, *domain.Question)
}

type Repositories struct {
	Decks     Decks
	Questions Questions
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Decks:     NewDecksRepo(db),
		Questions: NewQuestionsRepo(db),
	}
}
