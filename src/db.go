package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/logotipiwe/dc_go_utils/src/config"
	"log"
	"math/rand"
	"time"
)

var db *sql.DB

func InitDb() error {
	connectionStr := fmt.Sprintf("%v:%v@tcp(%v)/%v", GetConfig("DB_LOGIN"), GetConfig("DB_PASS"),
		GetConfig("DB_HOST"), GetConfig("DB_NAME"))
	conn, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return err
	}
	if err := conn.Ping(); err != nil {
		println(fmt.Sprintf("Error connecting database: %s", err))
		return err
	}
	db = conn
	println("Database connected!")
	return nil
}

func GetRandQuestionByLevel(level string) (error, *Question) {
	// Seed the random number generator
	rand.Int63n(time.Now().UnixNano())

	// Query to select a random row with the specified level
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
