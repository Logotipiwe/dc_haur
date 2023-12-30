package repo

import (
	"database/sql"
	"dc_haur/src/internal/domain"
	"github.com/jinzhu/gorm"
	"log"
)

type Decks interface {
	GetDecks() ([]domain.Deck, error)
}

type Questions interface {
	GetLevels(deckID string) ([]string, error)
	GetRandQuestion(deckID, levelName string) (*domain.Question, error)
	GetLevelsByName(deckName string) ([]string, error)
	GetRandQuestionByNames(deckName string, levelName string) (*domain.Question, error)
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
	gormDb, err := gorm.Open("mysql", db)
	if err != nil {
		log.Fatal(err)
	}
	return &Repositories{
		Decks:     NewDecksRepo(gormDb),
		Questions: NewQuestionsRepo(gormDb),
		History:   NewQuestionsHistoryRepo(gormDb),
	}
}
