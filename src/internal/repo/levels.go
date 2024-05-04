package repo

import (
	"dc_haur/src/internal/model"
	"github.com/jinzhu/gorm"
)

type Levels struct {
	db *gorm.DB
}

func NewLevelRepository(db *gorm.DB) *Levels {
	return &Levels{db: db}
}

func (r *Levels) GetByID(id string) (model.Level, error) {
	var level model.Level
	err := r.db.Raw("SELECT * FROM levels WHERE id = ?", id).First(&level).Error
	return level, err
}

func (r *Levels) GetLevelsNamesByDeckName(deckName string) ([]string, error) {
	var levels []model.Level
	err := r.db.Where("deck_id = (select id from decks where name = ?)", deckName).
		Order("level_order").Find(&levels).Error
	if err != nil {
		return nil, err
	}
	levelNames := make([]string, len(levels))
	for i, level := range levels {
		levelNames[i] = level.Name
	}
	return levelNames, nil
}

func (r *Levels) GetLevelsByDeckName(deckName string) ([]*model.Level, error) {
	var levels []*model.Level
	err := r.db.Where("deck_id = (select id from decks where name = ?)", deckName).Order("level_order").Find(&levels).Error
	return levels, err
}

func (r *Levels) GetLevelsByDeckId(deckID string) ([]*model.Level, error) {
	var levels []*model.Level
	err := r.db.Where("deck_id = ?", deckID).Order("level_order").Find(&levels).Error
	return levels, err
}

func (r *Levels) Create(level *model.Level) error {
	return r.db.Create(&level).Error
}

func (r *Levels) Update(level *model.Level) error {
	return r.db.Model(&level).Updates(level).Error
}

func (r *Levels) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&model.Level{}).Error
}

func (r *Levels) GetQuestionLevel(question *model.Question) (*model.Level, error) {
	var level model.Level
	err := r.db.Where("id = ?", question.LevelID).First(&level).Error
	return &level, err
}

func (r *Levels) GetQuestionsCount(id string) (int, error) {
	var count int
	err := r.db.Model(&model.Question{}).Where("level_id = ?", id).Count(&count).Error
	return count, err
}
func (r *Levels) GetOpenedQuestionsCount(id, clientID string) (int, error) {
	var count int
	sql := `SELECT COUNT(distinct question_id) FROM questions_history 
	   	WHERE client_id = ? 
		AND level_id = ?
		AND question_id in (select id from questions)` //если вопросы будут удаляться, чтобы это учлось
	err := r.db.Raw(sql, clientID, id).Count(&count).Error
	return count, err
}
