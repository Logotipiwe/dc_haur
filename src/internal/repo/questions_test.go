package repo

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetLevels_Success(t *testing.T) {
	// Create a new mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create a new QuestionsRepo with the mock database
	repo := NewQuestionsRepo(mockDB)

	// Define expected rows and columns
	rows := sqlmock.NewRows([]string{"level"}).
		AddRow("Easy").
		AddRow("Medium").
		AddRow("Hard")

	// Expect a SELECT query
	mock.ExpectQuery(`SELECT distinct q.level FROM questions q LEFT JOIN haur.decks d on d.id = q.deck_id WHERE d.name = ?`).
		WithArgs("ExampleDeck").
		WillReturnRows(rows)

	// Call the function
	err, levels := repo.GetLevels("ExampleDeck")

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the result
	expectedLevels := []string{"Easy", "Medium", "Hard"}
	if len(levels) != len(expectedLevels) {
		t.Errorf("Expected %d levels, got %d levels", len(expectedLevels), len(levels))
	}

	// Additional checks for each level can be added if necessary
}

func TestGetLevels_NoLevels(t *testing.T) {
	// Create a new mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create a new QuestionsRepo with the mock database
	repo := NewQuestionsRepo(mockDB)

	// Expect a SELECT query to return no rows
	mock.ExpectQuery(`SELECT distinct q.level FROM questions q LEFT JOIN haur.decks d on d.id = q.deck_id WHERE d.name = ?`).
		WithArgs("NoLevelsDeck").
		WillReturnRows(sqlmock.NewRows([]string{"level"}))

	// Call the function
	err, levels := repo.GetLevels("NoLevelsDeck")

	// Check for the expected error
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	// Ensure that levels is an empty slice in case of an error
	if levels != nil {
		t.Error("Expected levels to be nil")
	}
}

func TestGetRandQuestion_Error(t *testing.T) {
	// Create a new mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create a new QuestionsRepo with the mock database
	repo := NewQuestionsRepo(mockDB)

	// Expect a SELECT query to return an error
	mock.ExpectQuery(`SELECT q.id, q.level, q.deck_id, q.text FROM questions q LEFT JOIN decks d on d.id = q.deck_id WHERE level = ? AND d.name = ? ORDER BY rand() LIMIT 1`).
		WithArgs("Hard", "ErrorDeck").
		WillReturnError(errors.New("mocked error"))

	// Call the function
	err, question := repo.GetRandQuestion("ErrorDeck", "Hard")

	// Check for the expected error
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	// Ensure that question is nil in case of an error
	if question != nil {
		t.Error("Expected question to be nil, but it's not")
	}
}
