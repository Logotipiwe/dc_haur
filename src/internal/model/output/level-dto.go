package output

import "dc_haur/src/internal/model"

type LevelDto struct {
	model.Level
	Counts *QuestionsCounts `json:"counts,omitempty"`
}

type QuestionsCounts struct {
	QuestionsCount       int  `json:"questionsCount"`
	OpenedQuestionsCount *int `json:"openedQuestionsCount,omitempty"`
}
