package repo

import (
	"dc_haur/src/internal/domain"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"testing"
)

func TestGetDecks_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()
	gdb, err := gorm.Open("mysql", mockDB)
	if err != nil {
		t.Error(err)
	}
	repo := NewDecksRepo(gdb)
	rows := sqlmock.NewRows([]string{"id", "name", "description"}).
		AddRow(1, "Deck 1", "Description 1").
		AddRow(2, "Deck 2", "Description 2")
	mock.ExpectQuery(`SELECT id, name, description FROM decks`).WillReturnRows(rows)
	decks, err := repo.GetDecks()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedDecks := []domain.Deck{
		{ID: "1", Name: "Deck 1", Description: "Description 1"},
		{ID: "2", Name: "Deck 2", Description: "Description 2"},
	}

	if len(decks) != len(expectedDecks) {
		t.Errorf("Expected %d decks, got %d decks", len(expectedDecks), len(decks))
	}
}

func TestGetDecks_ErrorOnQuery(t *testing.T) {
	// Create a new mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()
	gdb, err := gorm.Open("mysql", mockDB)
	if err != nil {
		t.Error(err)
	}
	repo := NewDecksRepo(gdb)
	mock.ExpectQuery(`SELECT id, name, description FROM decks`).WillReturnError(errors.New("mocked error"))
	decks, err := repo.GetDecks()
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if decks != nil {
		t.Error("Expected decks to be nil, but it's not")
	}
}

func TestGetDecks_ErrorOnScanWithoutDescriptionField(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()
	gdb, err := gorm.Open("mysql", mockDB)
	if err != nil {
		t.Error(err)
	}
	repo := NewDecksRepo(gdb)
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Deck 1").
		AddRow(2, "Deck 2")
	mock.ExpectQuery(`SELECT id, name, description FROM decks`).WillReturnRows(rows)
	decks, err := repo.GetDecks()
	if err == nil {
		t.Error("Error expected")
	}
	if decks != nil {
		t.Errorf("Decks expected to be nil, but... %v", decks)
	}
}

func TestGetDecks_ErrorOnEmptyScan(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()
	gdb, err := gorm.Open("mysql", mockDB)
	if err != nil {
		t.Error(err)
	}
	repo := NewDecksRepo(gdb)
	mock.ExpectQuery(`SELECT id, name, description FROM decks`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}))
	decks, err := repo.GetDecks()
	if decks == nil {
		t.Error("Expected decks to be an empty slice, but it's nil")
	} else if len(decks) != 0 {
		t.Errorf("Expected decks to be an empty slice, but got %v", decks)
	}
}
