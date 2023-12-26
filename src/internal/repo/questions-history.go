package repo

import (
	"database/sql"
	"dc_haur/src/internal/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type QuestionsHistoryRepo struct {
	DB *sql.DB
}

func NewQuestionsHistoryRepo(DB *sql.DB) *QuestionsHistoryRepo {
	return &QuestionsHistoryRepo{DB: DB}
}

func (repo *QuestionsHistoryRepo) Insert(chatID int64, question *domain.Question) error {
	query := `INSERT INTO questions_history (id, deck_id, level_name, question_id, chat_id) VALUES (?, ?, ?, ?, ?)`
	_, err := repo.DB.Exec(query, uuid.NewString(), question.DeckID, question.Level, question.ID, chatID)
	return err
}
