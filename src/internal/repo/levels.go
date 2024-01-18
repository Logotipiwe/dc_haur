package repo

import (
	"dc_haur/src/internal/domain"
	"github.com/jinzhu/gorm"
)

type Levels struct {
	db *gorm.DB
}

func NewLevelRepository(db *gorm.DB) *Levels {
	return &Levels{db: db}
}

func (r *Levels) GetByID(id string) (domain.Level, error) {
	var level domain.Level
	err := r.db.Where("id = ?", id).First(&level).Error
	return level, err
}

func (r *Levels) GetLevelsNamesByDeckName(deckName string) ([]string, error) {
	var levels []domain.Level
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

func (r *Levels) GetLevelsByDeckName(deckName string) ([]*domain.Level, error) {
	var levels []*domain.Level
	err := r.db.Where("deck_id = (select id from decks where name = ?)", deckName).Order("level_order").Find(&levels).Error
	return levels, err
}

func (r *Levels) GetLevelsByDeckId(deckID string) ([]*domain.Level, error) {
	var levels []*domain.Level
	err := r.db.Where("deck_id = ?", deckID).Order("level_order").Find(&levels).Error
	return levels, err
}

func (r *Levels) Create(level *domain.Level) error {
	return r.db.Create(&level).Error
}

func (r *Levels) Update(level *domain.Level) error {
	return r.db.Model(&level).Updates(level).Error
}

func (r *Levels) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&domain.Level{}).Error
}

func (r *Levels) GetQuestionLevel(question *domain.Question) (*domain.Level, error) {
	var level domain.Level
	err := r.db.Where("id = ?", question.LevelID).First(&level).Error
	return &level, err
}
