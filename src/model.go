package main

type Question struct {
	ID     string `db:"id"`
	Level  string `db:"level"`
	DeckID string `db:"deck_id"`
	Text   string `db:"text"`
}

type Deck struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
