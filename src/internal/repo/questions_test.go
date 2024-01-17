package repo

import (
	"dc_haur/src/internal/domain"
	"errors"
	"github.com/jinzhu/gorm"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetLevels_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()
	gdb, err := gorm.Open("mysql", mockDB)
	if err != nil {
		t.Error(err)
	}
	repo := NewQuestionsRepo(gdb)
	rows := sqlmock.NewRows([]string{"level"}).AddRow("Easy").AddRow("Medium").AddRow("Hard")
	mock.ExpectQuery(regexp.QuoteMeta(GetLevelsByDeckNameSql)).
		WithArgs("ExampleDeck").
		WillReturnRows(rows)
	levels, err := repo.GetLevelsNames("ExampleDeck")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedLevels := []string{"Easy", "Medium", "Hard"}
	if len(levels) != len(expectedLevels) {
		t.Errorf("Expected %d levels, got %d levels", len(expectedLevels), len(levels))
	}
}

func TestGetLevels_NoLevels(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()
	gdb, err := gorm.Open("mysql", mockDB)
	if err != nil {
		t.Error(err)
	}
	repo := NewQuestionsRepo(gdb)
	mock.ExpectQuery(regexp.QuoteMeta(GetLevelsByDeckNameSql)).
		WithArgs("NoLevelsDeck").
		WillReturnRows(sqlmock.NewRows([]string{"level"}))
	levels, err := repo.GetLevelsNames("NoLevelsDeck")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if levels != nil {
		t.Error("Expected levels to be nil")
	}
}

func TestGetRandQuestion_Error(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()
	gdb, err := gorm.Open("mysql", mockDB)
	if err != nil {
		t.Error(err)
	}
	repo := NewQuestionsRepo(gdb)
	mock.ExpectQuery(regexp.QuoteMeta(GetRandQuestionByDeckNameSql)).
		WithArgs("Hard", "ErrorDeck").
		WillReturnError(errors.New("mocked error"))
	question, err := repo.GetRandQuestionByNames("ErrorDeck", "Hard")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if question != nil {
		t.Error("Expected question to be nil, but it's not")
	}
}

func TestGetRandQuestion_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()
	gdb, err := gorm.Open("mysql", mockDB)
	if err != nil {
		t.Error(err)
	}
	repo := NewQuestionsRepo(gdb)
	rows := sqlmock.NewRows([]string{"id", "level", "deck_id", "text"}).
		AddRow("1", "Hard", "2", "someRandText")
	mock.ExpectQuery(regexp.QuoteMeta(GetRandQuestionByDeckNameSql)).
		WithArgs("Hard", "ExampleDeck").
		WillReturnRows(rows)
	question, err := repo.GetRandQuestionByNames("ExampleDeck", "Hard")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if question == nil {
		t.Errorf("Question is nil")
	}
	expectedQuestion := domain.Question{
		ID:     "1",
		Level:  "Hard",
		DeckID: "2",
		Text:   "someRandText",
	}
	if *question != expectedQuestion {
		t.Errorf("Question %v not equals to expected %v", *question, expectedQuestion)
	}
}

func TestGetLevels_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open a stub database connection: %v", err)
	}
	defer db.Close()

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		t.Error(err)
	}
	questionsRepo := NewQuestionsRepo(gdb)

	mock.ExpectQuery(regexp.QuoteMeta(GetLevelsByDeckNameSql)).
		WithArgs("testDeck").
		WillReturnError(errors.New("database error"))

	_, err = questionsRepo.GetLevelsNames("testDeck")
	if err == nil || err.Error() != "database error" {
		t.Errorf("Expected database error, but got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
