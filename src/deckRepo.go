package main

func GetDecks() (error, []Deck) {
	sql := `
		SELECT id, name, description FROM decks
`
	res, err := db.Query(sql)
	if err != nil {
		return err, nil
	}
	var decks = make([]Deck, 0)
	for res.Next() {
		deck := Deck{}
		err := res.Scan(&deck.ID, &deck.Name, &deck.Description)
		if err != nil {
			return err, nil
		}
		decks = append(decks, deck)
	}
	return nil, decks
}
