package main

import (
	errors "errors"
	"log"
)

func GetLevels(deckName string) (error, []string) {
	var ans []string
	res, err := db.Query("SELECT distinct q.level FROM questions q LEFT JOIN haur.decks d on d.id = q.deck_id WHERE d.name = ?", deckName)
	if err != nil {
		return err, nil
	}
	defer res.Close()
	for res.Next() {
		var level string
		err := res.Scan(&level)
		if err != nil {
			return err, nil
		}
		ans = append(ans, level)
	}
	if len(ans) == 0 {
		return errors.New("NoLevels from deck " + deckName), nil
	}
	return nil, ans
}

func GetRandQuestion(deckName string, levelName string) (error, *Question) {
	query := `
		SELECT q.id, q.level, q.deck_id, q.text
		FROM questions q
		LEFT JOIN decks d on d.id = q.deck_id
		WHERE level = ? AND d.name = ?
		ORDER BY rand()
		LIMIT 1`

	var result Question
	err := db.QueryRow(query, levelName, deckName).Scan(&result.ID, &result.Level, &result.DeckID, &result.Text)
	if err != nil {
		log.Println("Error getting question from DB.")
		log.Println(err)
		return err, nil
	}
	return nil, &result
}
