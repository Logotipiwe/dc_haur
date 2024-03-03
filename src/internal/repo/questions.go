package repo

import (
	"dc_haur/src/internal/domain"
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
	GetLevelsSql       = "SELECT name FROM levels WHERE deck_id = ? ORDER BY level_order"
	GetRandQuestionSql = "SELECT q.id, q.level_id, q.text FROM (select q.*, (select count(*) from questions_history where question_id = q.id) asked from questions q) q  WHERE q.level_id = ? ORDER BY q.asked, rand() LIMIT 1"
)

var (
	NoLevelsErr = errors.New("no levels from deck")
)

func (r *Questions) GetRandQuestion(levelID string) (*domain.Question, error) {
	var result domain.Question
	err := r.db.Raw(GetRandQuestionSql, levelID).Row().Scan(&result.ID, &result.LevelID, &result.Text)
	if err != nil {
		log.Println("Error getting question from DB.")
		log.Println(err)
		return nil, err
	}
	return &result, nil
}

func (r *Questions) GetRandQuestionByNames(deckName string, levelNameWithEmoji string) (*domain.Question, error) {
	var level domain.Level
	if err := r.db.Where("(concat(coalesce(concat(emoji, ' '),''), name) = ?) AND deck_id = (select id from decks where name = ?)",
		levelNameWithEmoji, deckName).
		First(&level).Error; err != nil {
		return nil, err
	}
	return r.GetRandQuestion(level.ID)
}
