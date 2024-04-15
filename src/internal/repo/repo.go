package repo

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"log"
)

//type Decks interface {
//	GetDecks() ([]model.Deck, error)
//}
//
//type Questions interface {
//	GetLevels(deckID string) ([]*model.Level, error)
//	GetLevelsNames(deckID string) ([]string, error)
//	GetRandQuestion(levelID string) (*model.Question, error)
//	GetLevelsByName(deckName string) ([]string, error)
//	GetRandQuestionByNames(deckName string, levelName string) (*model.Question, error)
//}
//
//type History interface {
//	Insert(string, *model.Question) error
//	Truncate() error
//}

type Repositories struct {
	*Decks
	*Questions
	*History
	*Levels
	*VectorImages
	*DeckLikes
	*QuestionLikes
	*UsedQuestions
}

func NewRepositories(db *sql.DB) *Repositories {
	gormDb, err := NewGorm(db)
	if err != nil {
		log.Fatal(err)
	}
	return &Repositories{
		Decks:         NewDecksRepo(gormDb),
		Questions:     NewQuestionsRepo(gormDb),
		History:       NewQuestionsHistoryRepo(gormDb),
		Levels:        NewLevelRepository(gormDb),
		VectorImages:  NewVectorImages(gormDb),
		DeckLikes:     NewDeckLikes(gormDb),
		QuestionLikes: NewQuestionLikes(gormDb),
		UsedQuestions: NewUsedQuestionsRepo(gormDb),
	}
}

func NewGorm(db *sql.DB) (*gorm.DB, error) {
	return gorm.Open("mysql", db)
}
