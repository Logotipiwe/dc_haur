package repo

import (
	"dc_haur/src/internal/model"
	"github.com/jinzhu/gorm"
)

type UsedQuestions struct {
	DB *gorm.DB
}

func NewUsedQuestionsRepo(DB *gorm.DB) *UsedQuestions {
	return &UsedQuestions{DB: DB}
}

func (repo *UsedQuestions) Insert(clientID string, question *model.Question) error {
	tableName := model.UsedQuestion{}.TableName()
	return repo.DB.Exec(`INSERT INTO `+tableName+` (question_id, client_id) VALUES (?,?) on duplicate key update question_id=question_id`,
		question.ID, clientID).Error
}

func (repo *UsedQuestions) Truncate() error {
	return repo.DB.Exec(`TRUNCATE TABLE used_questions`).Error
}

func (repo *UsedQuestions) AreAllQuestionsUsed(levelID string, clientID string) (allUsed bool, err error) {
	countOfUsed := "(select count(*) from used_questions where question_id in (select id from questions where level_id = ?) and client_id = ?)"
	countInLevel := "(select count(*) from questions where level_id = ?)"
	err = repo.DB.Raw("SELECT used >= in_level from (select "+countOfUsed+" used, "+countInLevel+" in_level) as numbers", levelID, clientID, levelID).Row().Scan(&allUsed)
	if err != nil {
		return false, err
	}
	return allUsed, nil
}

func (repo *UsedQuestions) Clear(levelID string, clientId string) error {
	return repo.DB.Exec("DELETE FROM used_questions WHERE question_id in (select id from questions where level_id = ?) AND client_id = ?", levelID, clientId).Error
}
