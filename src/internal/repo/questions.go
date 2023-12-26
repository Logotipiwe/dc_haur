package repo

import (
	"dc_haur/src/internal/domain"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

type QuestionsRepo struct {
	db *gorm.DB
}

func NewQuestionsRepo(db *gorm.DB) *QuestionsRepo {
	return &QuestionsRepo{db: db}
}

const GetLevelsSql = "SELECT distinct q.level FROM questions q LEFT JOIN haur.decks d on d.id = q.deck_id WHERE d.name = ?"
const GetRandQuestionSql = "SELECT q.id, q.level, q.deck_id, q.text FROM (select q.*, (select count(*) from questions_history where question_id = q.id) asked from questions q) q LEFT JOIN decks d on d.id = q.deck_id WHERE level = ? AND d.name = ? ORDER BY q.asked, rand() LIMIT 1"

func (r *QuestionsRepo) GetLevels(deckName string) ([]string, error) {
	var ans []string
	rows, err := r.db.Raw(GetLevelsSql, deckName).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var level string
	for rows.Next() {
		rows.Scan(&level)
		ans = append(ans, level)
	}
	if len(ans) == 0 {
		return nil, errors.New("NoLevels from deck " + deckName)
	}
	return ans, nil
}

func (r *QuestionsRepo) GetRandQuestion(deckName string, levelName string) (*domain.Question, error) {
	var result domain.Question
	err := r.db.Raw(GetRandQuestionSql, levelName, deckName).Row().Scan(&result.ID, &result.Level, &result.DeckID, &result.Text)
	if err != nil {
		log.Println("Error getting question from DB.")
		log.Println(err)
		return nil, err
	}
	return &result, nil
}
