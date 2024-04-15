package repo

import (
	"dc_haur/src/internal/model"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

type Questions struct {
	db *gorm.DB
}

func NewQuestionsRepo(db *gorm.DB) *Questions {
	return &Questions{db: db}
}

const (
	GetRandQuestionSql     = "SELECT q.id, q.level_id, q.text, q.additional_text FROM questions q where q.level_id = ? AND id not in (select question_id from used_questions where client_id = ?) ORDER BY rand() LIMIT 1"
	GetRandQuestionSqlDumb = "SELECT q.id, q.level_id, q.text, q.additional_text FROM questions q where q.level_id = ? ORDER BY rand() LIMIT 1"
)

var (
	NoLevelsErr = errors.New("no levels from deck")
)

func (r *Questions) GetRandQuestionByLevel(levelID string, clientId string) (*model.Question, error) {
	var result model.Question
	err := r.db.Raw(GetRandQuestionSql, levelID, clientId).Row().Scan(&result.ID, &result.LevelID, &result.Text, &result.AdditionalText)
	if err != nil {
		log.Println("Error getting question from DB.")
		log.Println(err)
		return nil, err
	}
	return &result, nil
}

func (r *Questions) GetRandQuestionDumb(levelID string) (*model.Question, error) {
	var result model.Question
	err := r.db.Raw(GetRandQuestionSqlDumb, levelID).Row().Scan(&result.ID, &result.LevelID, &result.Text, &result.AdditionalText)
	if err != nil {
		log.Println("Error getting question from DB.")
		log.Println(err)
		return nil, err
	}
	return &result, nil
}

func (r *Questions) GetLevelByNames(deckName string, levelNameWithEmoji string) (*model.Level, error) {
	var level model.Level
	err := r.db.Where("(concat(coalesce(concat(emoji, ' '),''), name) = ?) AND deck_id = (select id from decks where name = ?)",
		levelNameWithEmoji, deckName).First(&level).Error
	return &level, err
}

func (r *Questions) GetAllByDeckId(deckId string) ([]model.Question, error) {
	questions := make([]model.Question, 0)
	if err := r.db.Find(&questions, "level_id in (select id from levels where deck_id = ?)", deckId).Error; err != nil {
		return nil, err
	}
	return questions, nil
}
