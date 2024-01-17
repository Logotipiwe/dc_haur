package repo

import (
	"dc_haur/src/internal/domain"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type History struct {
	DB *gorm.DB
}

func NewQuestionsHistoryRepo(DB *gorm.DB) *History {
	return &History{DB: DB}
}

func (repo *History) Insert(chatID string, question *domain.Question) error {
	query := &domain.QuestionHistory{
		ID:         uuid.NewString(),
		LevelID:    question.LevelID,
		QuestionID: question.ID,
		ChatID:     chatID,
	}
	return repo.DB.Create(query).Error
}

func (repo *History) Truncate() error {
	return repo.DB.Exec(`TRUNCATE TABLE questions_history`).Error
}
