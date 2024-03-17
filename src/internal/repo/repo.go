package repo

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"log"
)

//type Decks interface {
//	GetDecks() ([]domain.Deck, error)
//}
//
//type Questions interface {
//	GetLevels(deckID string) ([]*domain.Level, error)
//	GetLevelsNames(deckID string) ([]string, error)
//	GetRandQuestion(levelID string) (*domain.Question, error)
//	GetLevelsByName(deckName string) ([]string, error)
//	GetRandQuestionByNames(deckName string, levelName string) (*domain.Question, error)
//}
//
//type History interface {
//	Insert(string, *domain.Question) error
//	Truncate() error
//}

type Repositories struct {
	*Decks
	*Questions
	*History
	*Levels
	*VectorImages
}

func NewRepositories(db *sql.DB) *Repositories {
	gormDb, err := NewGorm(db)
	if err != nil {
		log.Fatal(err)
	}
	return &Repositories{
		Decks:        NewDecksRepo(gormDb),
		Questions:    NewQuestionsRepo(gormDb),
		History:      NewQuestionsHistoryRepo(gormDb),
		Levels:       NewLevelRepository(gormDb),
		VectorImages: NewVectorImages(gormDb),
	}
}

func NewGorm(db *sql.DB) (*gorm.DB, error) {
	return gorm.Open("mysql", db)
}
