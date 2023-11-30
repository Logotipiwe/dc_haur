package main

const (
	EASY   = "EASY"
	MEDIUM = "MEDIUM"
	HARD   = "HARD"
)

var QuestionsLevels = map[string]string{
	EASY:   "Знакомство",
	MEDIUM: "Погружение",
	HARD:   "Рефлексия",
}

var QuestionsLevelsReverted = map[string]string{
	"Знакомство": EASY,
	"Погружение": MEDIUM,
	"Рефлексия":  HARD,
}

type Question struct {
	ID     string `db:"id"`
	Level  string `db:"level"`
	DeckID string `db:"deck_id"`
	Text   string `db:"text"`
}
