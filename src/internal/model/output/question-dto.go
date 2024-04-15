package output

import "dc_haur/src/internal/model"

type QuestionDTO struct {
	model.Question
	IsLast bool `json:"isLast"`
}
