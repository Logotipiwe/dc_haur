package tests

import (
	"database/sql"
	"dc_haur/src/internal/mocks"
	"dc_haur/src/internal/repo"
	"dc_haur/src/internal/service"
	utils "dc_haur/src/pkg"
	"dc_haur/src/tghttp"
	"github.com/DATA-DOG/go-sqlmock"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"regexp"
	"testing"
)

type MessageTestEntry struct {
	Update      tgbotapi.Update
	WantReplies []string
}

func NewMessageTestEntry(message string, wantReplies ...string) *MessageTestEntry {
	return &MessageTestEntry{
		Update: tgbotapi.Update{
			Message: &tgbotapi.Message{
				Text: message,
				Chat: &tgbotapi.Chat{ID: int64(4356745)},
				From: &tgbotapi.User{UserName: "@Logotipiwe"}}},
		WantReplies: wantReplies}
}

func TestIntegration(t *testing.T) {
	db := setupDbMock(t)
	repoMocks := mocks.NewRepoMocksWithDb(db)
	servicesMocks := service.NewServices(repoMocks.QuestionRepo, repoMocks.DeckRepo, mocks.NewBotInteractorMock())
	handler := service.NewHandler(servicesMocks.TgMessages, servicesMocks.Cache)
	t.Run("GetRandQuestionNormally", func(t *testing.T) {
		testSequentialUpdates(t, handler, "start", []*MessageTestEntry{
			NewMessageTestEntry("/start", service.WelcomeMessage),
			NewMessageTestEntry("Deck 1", service.GotLevelsMessage),
			NewMessageTestEntry("level 11", "q111"),
		})
	})
	t.Run("GetRandQuestionFromNonExistingLevel", func(t *testing.T) {
		testSequentialUpdates(t, handler, "start", []*MessageTestEntry{
			NewMessageTestEntry("/start", service.WelcomeMessage),
			NewMessageTestEntry("Deck 1", service.GotLevelsMessage),
			NewMessageTestEntry("level 21", tghttp.ErrorOrUnknownMessage),
		})
	})
}

func setupDbMock(t *testing.T) *sql.DB {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %s", err)
	}
	mock.MatchExpectationsInOrder(false)
	data := []struct {
		ID          string
		Name        string
		Description string
		Levels      []struct {
			Name      string
			Questions []string
		}
	}{
		{ID: "1", Name: "Deck 1", Description: "Desc 1", Levels: []struct {
			Name      string
			Questions []string
		}{{Name: "level 11", Questions: []string{"q111", "q112"}},
			{Name: "level 12", Questions: []string{"q121", "q122"}}}},
		{ID: "2", Name: "Deck 2", Description: "Desc 2", Levels: []struct {
			Name      string
			Questions []string
		}{{Name: "level 21", Questions: []string{"q211", "q212"}},
			{Name: "level 22", Questions: []string{"q221", "q222"}},
			{Name: "level 23", Questions: []string{"q231", "q232"}}}},
		{ID: "3", Name: "Deck 3", Description: "Desc 3", Levels: []struct {
			Name      string
			Questions []string
		}{{Name: "level 31", Questions: []string{"q311", "q312"}},
			{Name: "level 32", Questions: []string{"q321", "q322"}}}},
	}
	deckRows := sqlmock.NewRows([]string{"id", "name", "description"})
	for _, deck := range data {
		deckRows.AddRow(deck.ID, deck.Name, deck.Description)
	}
	mock.ExpectQuery(repo.GetDecksQuery).WillReturnRows(deckRows)

	for _, deck := range data {
		levelsRows := sqlmock.NewRows([]string{"level"})
		for _, level := range deck.Levels {
			levelsRows.AddRow(level.Name)
		}
		mock.ExpectQuery(repo.GetLevelsSql).WithArgs(deck.Name).WillReturnRows(levelsRows)
	}

	for _, deck := range data {
		for _, level := range deck.Levels {
			questionRow := sqlmock.NewRows([]string{"id", "level", "deck_id", "text"})
			questionRow.AddRow("1", level.Name, deck.ID, level.Questions[rand.Intn(len(level.Questions))])
			mock.ExpectQuery(regexp.QuoteMeta(repo.GetRandQuestionSql)).WithArgs(level.Name, deck.Name).WillReturnRows(questionRow)
		}
	}
	return db
}

func testSequentialUpdates(t *testing.T, handler *service.Handler, name string, tests []*MessageTestEntry) {
	t.Run(name, func(t *testing.T) {
		for _, messageTestEntry := range tests {
			reply, err := handler.HandleMessageAndReply(messageTestEntry.Update)
			if err != nil {
				t.Errorf("error returned from HandleMessageAndReply: %v", err)
			} else {
				if replyText, ok := reply.(*tgbotapi.MessageConfig); ok {
					if !utils.ExistsInArr(replyText.Text, messageTestEntry.WantReplies) {
						t.Errorf("Unexpected reply text. Expected one of: %v. Got: %s",
							messageTestEntry.WantReplies, replyText.Text)
					}
				} else if _, ok := reply.(*tgbotapi.PhotoConfig); ok {
					//it's ok)
				} else {
					t.Errorf("unknown reply type")
				}
			}
		}
	})
}
