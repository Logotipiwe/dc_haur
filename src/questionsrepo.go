package main

import (
	errors "errors"
	"log"
)

func GetLevels( /*deckID*/ ) (error, []string) {
	deckID := "1"
	var ans []string
	res, err := db.Query("SELECT distinct level FROM questions WHERE deck_id = ?", deckID)
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
		return errors.New("NoLevels from deck " + deckID), nil
	}
	return nil, ans
}

func GetRandQuestionByLevel(level string) (error, *Question) {
	query := `
		SELECT id, level, deck_id, text
		FROM questions
		WHERE level = ?
		ORDER BY rand()
		LIMIT 1`

	var result Question
	err := db.QueryRow(query, level).Scan(&result.ID, &result.Level, &result.DeckID, &result.Text)
	if err != nil {
		log.Println("Error getting question from DB.")
		log.Println(err)
		return err, nil
	}
	return nil, &result
}
