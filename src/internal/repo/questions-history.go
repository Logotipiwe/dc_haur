package repo

import (
	"dc_haur/src/internal/domain"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type QuestionsHistoryRepo struct {
	DB *gorm.DB
}

func NewQuestionsHistoryRepo(DB *gorm.DB) *QuestionsHistoryRepo {
	return &QuestionsHistoryRepo{DB: DB}
}

func (repo *QuestionsHistoryRepo) Insert(chatID string, question *domain.Question) error {
	query := &domain.QuestionHistory{
		ID:         uuid.NewString(),
		DeckID:     question.DeckID,
		LevelName:  question.Level,
		QuestionID: question.ID,
		ChatID:     chatID,
	}
	return repo.DB.Create(query).Error
}

func (repo *QuestionsHistoryRepo) Truncate() error {
	return repo.DB.Exec(`TRUNCATE TABLE questions_history`).Error
}
