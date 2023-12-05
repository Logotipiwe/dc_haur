package repo

import (
	"database/sql"
	"dc_haur/src/internal/domain"
	errors "errors"
	"log"
)

type QuestionsRepo struct {
	db *sql.DB
}

func NewQuestionsRepo(db *sql.DB) *QuestionsRepo {
	return &QuestionsRepo{db: db}
}

func (r *QuestionsRepo) GetLevels(deckName string) (error, []string) {
	var ans []string
	res, err := r.db.Query("SELECT distinct q.level FROM questions q LEFT JOIN haur.decks d on d.id = q.deck_id WHERE d.name = ?", deckName)
	if err != nil {
		return err, nil
	}
	defer res.Close()
	for res.Next() {
		var level string
		err := res.Scan(&level)
		if err != nil {
			return err, nil
		}
		ans = append(ans, level)
	}
	if len(ans) == 0 {
		return errors.New("NoLevels from deck " + deckName), nil
	}
	return nil, ans
}

func (r *QuestionsRepo) GetRandQuestion(deckName string, levelName string) (error, *domain.Question) {
	query := "SELECT q.id, q.level, q.deck_id, q.text FROM questions q LEFT JOIN decks d on d.id = q.deck_id WHERE level = ? AND d.name = ? ORDER BY rand() LIMIT 1"

	var result domain.Question
	err := r.db.QueryRow(query, levelName, deckName).Scan(&result.ID, &result.Level, &result.DeckID, &result.Text)
	if err != nil {
		log.Println("Error getting question from DB.")
		log.Println(err)
		return err, nil
	}
	return nil, &result
}
