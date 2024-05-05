package tests

import (
	"bytes"
	http2 "dc_haur/src/http"
	"dc_haur/src/internal/model"
	"dc_haur/src/internal/model/output"
	"dc_haur/src/internal/service"
	utils "dc_haur/src/pkg"
	"encoding/json"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	config "github.com/logotipiwe/dc_go_config_lib"
	"github.com/stretchr/testify/assert"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

const (
	d1   = "em1 deck d1 name"
	d2   = "em2 deck d2 name"
	d3   = "deck d3 name"
	d4   = "deck d4 name"
	d1l1 = "em1 l1"
	d1l2 = "em1 l2"
	d1l3 = "l3"
	d2l1 = "em2 l1"
	d2l2 = "l2"
	d2l3 = "em2 l3"
	d3l1 = "em3 l1"
	d3l2 = "l2"
)

const clientID = "integrationTestsClient"

func TestApplication(t *testing.T) {

	res, err := strconv.ParseBool(config.GetConfigOr("DO_INTEGRATION_TESTS", "true"))
	if err == nil && !res {
		//couldn't find way to not execute it in unit tests run
		println("TEST SKIPPED")
		return
	}

	checkIfImagesEnabled(t)

	/*DECLARATIONS*/
	em11 := "em1"
	expectedD1L1 := model.Level{
		ID:          "d1l1",
		DeckID:      "d1",
		LevelOrder:  1,
		Name:        "l1",
		Emoji:       &em11,
		ColorStart:  "0,0,0",
		ColorEnd:    "255,255,255",
		ColorButton: "1,1,1",
	}

	expectedD1L2 := model.Level{
		ID:          "d1l2",
		DeckID:      "d1",
		LevelOrder:  2,
		Name:        "l2",
		Emoji:       &em11,
		ColorStart:  "0,0,0",
		ColorEnd:    "255,255,255",
		ColorButton: "2,2,2",
	}
	expectedD1L3 := model.Level{
		ID:          "d1l3",
		DeckID:      "d1",
		LevelOrder:  3,
		Name:        "l3",
		Emoji:       nil,
		ColorStart:  "0,0,0",
		ColorEnd:    "255,255,255",
		ColorButton: "3,3,3",
	}
	/*END DECLARATIONS*/

	checkTg, err := strconv.ParseBool(config.GetConfigOr("CHECK_TG", "true"))
	assert.NoError(t, err)
	if checkTg {
		t.Run("Telegram client", func(t *testing.T) {
			t.Run("start message", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject("/start")
				ans := sendUpdate(t, update)
				assert.Equal(t, service.WelcomeMessage, ans.Text)
			})

			t.Run("get decks start", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject("/start")
				ans := sendUpdate(t, update)
				replyMarkup := toMarkup(t, ans.BaseChat.ReplyMarkup)
				assert.Equal(t, 4, len(replyMarkup.Keyboard))
				assert.Equal(t, d1, replyMarkup.Keyboard[0][0].Text)
				assert.Equal(t, d2, replyMarkup.Keyboard[1][0].Text)
				assert.Equal(t, d3, replyMarkup.Keyboard[2][0].Text)
				assert.Equal(t, d4, replyMarkup.Keyboard[3][0].Text)
				println(ans)
			})

			t.Run("get decks start from group", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject("/start@HowAreYouReallyGameBot")
				ans := sendUpdate(t, update)
				replyMarkup := toMarkup(t, ans.BaseChat.ReplyMarkup)
				assert.Equal(t, 4, len(replyMarkup.Keyboard))
				assert.Equal(t, d1, replyMarkup.Keyboard[0][0].Text)
				assert.Equal(t, d2, replyMarkup.Keyboard[1][0].Text)
				assert.Equal(t, d3, replyMarkup.Keyboard[2][0].Text)
				assert.Equal(t, d4, replyMarkup.Keyboard[3][0].Text)
				println(ans)
			})

			t.Run("select deck", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject("/start")
				ans := sendUpdate(t, update)
				update = createUpdateObject(d1)
				ans = sendUpdate(t, update)
				replyMarkup := toMarkup(t, ans.BaseChat.ReplyMarkup)
				//TODO check deck description here
				assert.True(t, strings.HasSuffix(ans.Text, service.GotLevelsMessage))
				assert.Equal(t, 3, len(replyMarkup.Keyboard[0]))
				assert.Equal(t, d1l1, replyMarkup.Keyboard[0][0].Text)
				assert.Equal(t, d1l2, replyMarkup.Keyboard[0][1].Text)
				assert.Equal(t, d1l3, replyMarkup.Keyboard[0][2].Text)
				println(ans)
			})

			t.Run("select deck with no emoji", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject("/start")
				ans := sendUpdate(t, update)
				update = createUpdateObject(d3)
				ans = sendUpdate(t, update)
				replyMarkup := toMarkup(t, ans.BaseChat.ReplyMarkup)
				//TODO check deck description here
				assert.True(t, strings.HasSuffix(ans.Text, service.GotLevelsMessage))
				assert.Equal(t, 2, len(replyMarkup.Keyboard[0]))
				assert.Equal(t, d3l1, replyMarkup.Keyboard[0][0].Text)
				assert.Equal(t, d3l2, replyMarkup.Keyboard[0][1].Text)
				println(ans)
			})

			t.Run("select deck; select level", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject("/start")
				ans := sendUpdate(t, update)
				update = createUpdateObjectFrom(update, d1)
				ans = sendUpdate(t, update)
				update = createUpdateObjectFrom(update, d1l1)
				ans = sendUpdate(t, update)
				assert.Contains(t, []string{"question d1l1q1 text", "question d1l1q2 text", "question d1l1q3 text"}, ans.Text)
				println(ans)
			})

			t.Run("select level > markup nil", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject("/start")
				ans := sendUpdate(t, update)
				update = createUpdateObjectFrom(update, d1)
				ans = sendUpdate(t, update)
				update = createUpdateObjectFrom(update, d1l1)
				ans = sendUpdate(t, update)
				assert.Nil(t, ans.BaseChat.ReplyMarkup)
				println(ans)
			})

			t.Run("select deck; select level many times", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject("/start")
				ans := sendUpdate(t, update)
				update = createUpdateObjectFrom(update, d1)
				ans = sendUpdate(t, update)
				for i := 0; i < 10; i++ {
					update = createUpdateObjectFrom(update, d1l1)
					ans = sendUpdate(t, update)
					assert.Contains(t, []string{"question d1l1q1 text", "question d1l1q2 text", "question d1l1q3 text"}, ans.Text)
					assert.Nil(t, ans.BaseChat.ReplyMarkup)
				}
			})

			t.Run("questions in level are ordered", func(t *testing.T) {
				defer failOnPanic(t)
				clearHistory(t)
				questions := []string{"question d1l1q1 text", "question d1l1q2 text", "question d1l1q3 text"}
				update := createUpdateObject("/start")
				ans := sendUpdate(t, update)
				update = createUpdateObjectFrom(update, "deck d1 name")
				ans = sendUpdate(t, update)
				for i := 0; i < 5; i++ {
					update = createUpdateObjectFrom(update, d1l1)
					ans = sendUpdate(t, update)
					ansIndex1 := utils.FindIndex(questions, ans.Text)
					assert.NotEqual(t, -1, ansIndex1)

					update = createUpdateObjectFrom(update, d1l1)
					ans = sendUpdate(t, update)
					ansIndex2 := utils.FindIndex(questions, ans.Text)
					assert.NotEqual(t, -1, ansIndex2)
					assert.NotEqual(t, ansIndex1, ansIndex2)

					update = createUpdateObjectFrom(update, d1l1)
					ans = sendUpdate(t, update)
					ansIndex3 := utils.FindIndex(questions, ans.Text)
					assert.NotEqual(t, -1, ansIndex3)
					assert.NotEqual(t, ansIndex1, ansIndex3)
					assert.NotEqual(t, ansIndex2, ansIndex3)
					println("ORDER CHECK FINISHED")
					time.Sleep(100 * time.Millisecond)
				}
			})

			t.Run("/question command", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject(service.QuestionCommand)
				ans := sendUpdate(t, update)
				assert.Equal(t, ans.Text, service.AssignNewQuestionText)
				update = createUpdateObjectFrom(update, "what??")
				ans = sendUpdate(t, update)
				assert.Equal(t, ans.Text, service.AcceptNewQuestionText)
			})

			t.Run("/question command", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject(service.QuestionCommand + "@HowAreYouReallyGameBot")
				ans := sendUpdate(t, update)
				assert.Equal(t, ans.Text, service.AssignNewQuestionText)
				update = createUpdateObjectFrom(update, "what??")
				ans = sendUpdate(t, update)
				assert.Equal(t, ans.Text, service.AcceptNewQuestionText)
			})

			t.Run("/feedback command from group", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject(service.FeedbackCommand)
				ans := sendUpdate(t, update)
				assert.Equal(t, ans.Text, service.AssignFeedbackText)
				update = createUpdateObjectFrom(update, "MyFeedback")
				ans = sendUpdate(t, update)
				assert.Equal(t, ans.Text, service.AcceptFeedbackText)
			})

			t.Run("/feedback command from group", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject(service.FeedbackCommand + "@HowAreYouReallyGameBot")
				ans := sendUpdate(t, update)
				assert.Equal(t, ans.Text, service.AssignFeedbackText)
				update = createUpdateObjectFrom(update, "MyFeedback")
				ans = sendUpdate(t, update)
				assert.Equal(t, ans.Text, service.AcceptFeedbackText)
			})
		})
	}

	checkWeb, err := strconv.ParseBool(config.GetConfigOr("CHECK_WEB", "true"))
	assert.NoError(t, err)
	if checkWeb {
		t.Run("Api client", func(t *testing.T) {
			appUrl := config.GetConfig("TEST_URL")
			println("Url to test " + appUrl)
			apiV1 := "/api/v1"
			//apiV3 := "/api/v3"
			apiIntegrationTest := apiV1 + "/integration-test"

			t.Run("/test-image secured", func(t *testing.T) {
				code, err := getResponseCode("GET", appUrl+apiIntegrationTest+"/test-image")
				assert.Nil(t, err)
				assert.Equal(t, 401, code)
			})
			t.Run("/test-chat secured", func(t *testing.T) {
				code, err := getResponseCode("POST", appUrl+apiIntegrationTest+"/test-chat")
				assert.Nil(t, err)
				assert.Equal(t, 401, code)
			})
			t.Run("/clear-history secured", func(t *testing.T) {
				code, err := getResponseCode("POST", appUrl+apiIntegrationTest+"/clear-history")
				assert.Nil(t, err)
				assert.Equal(t, 401, code)
			})
			t.Run("/images-enabled secured", func(t *testing.T) {
				code, err := getResponseCode("GET", appUrl+apiIntegrationTest+"/images-enabled")
				assert.Nil(t, err)
				assert.Equal(t, 401, code)
			})

			t.Run("get decks by language", func(t *testing.T) {
				defer failOnPanic(t)

				resultRu := getDecksFromApiV3(t, "RU", clientID)
				resultEn := getDecksFromApiV3(t, "EN", clientID)

				assert.Equal(t, 2, len(resultEn))
				assert.Equal(t, 1, len(resultRu))

				result := make([]output.DeckDTO, 3)
				result = append(result, resultRu[0], resultEn[0], resultEn[1])

				for i := range result {
					assert.NotNil(t, result[i].ID)
					assert.NotNil(t, result[i].LanguageCode)
					assert.NotNil(t, result[i].Name)
					assert.NotNil(t, result[i].Description)
					assert.NotNil(t, result[i].Labels)
					assert.NotNil(t, result[i].ImageID)
					assert.NotNil(t, result[i].CardsCount)
				}
			})

			t.Run("check decks cards count field", func(t *testing.T) {
				defer failOnPanic(t)

				result := getDecksFromApiV3(t, "EN", clientID)
				assert.NotEmpty(t, result)

				expectedCounts := []int{8, 3, 3}
				for i := range result {
					assert.NotNil(t, result[i].CardsCount)
					assert.Equal(t, expectedCounts[i], result[i].CardsCount)
				}
			})

			t.Run("check decks opened cards count", func(t *testing.T) {
				defer failOnPanic(t)

				//deck 1
				getQuestionFromApi(t, "d1l1", clientID, appUrl+apiV1)
				getQuestionFromApi(t, "d1l1", clientID, appUrl+apiV1)
				getQuestionFromApi(t, "d1l2", clientID, appUrl+apiV1)

				//deck 2 (only one card of this level, so should be 1 opened)
				getQuestionFromApi(t, "d2l1", clientID, appUrl+apiV1)
				getQuestionFromApi(t, "d2l1", clientID, appUrl+apiV1)

				result := getDecksFromApiV3(t, "EN", clientID)

				expectedCounts := []int{3, 1, 0}

				assert.NotEmpty(t, result)

				for i := range result {
					assert.NotNil(t, result[i].OpenedCount)
					assert.Equal(t, expectedCounts[i], result[i].OpenedCount)
				}
			})

			t.Run("get localized available decks", func(t *testing.T) {
				defer failOnPanic(t)

				languageCode := "EN"
				result := getDecksFromApiV3(t, languageCode, clientID)

				assert.Equal(t, 2, len(result))
				for i := range result {
					checkDeckDTOFields(t, result[i])
				}

				languageCode = "RU"
				result = getDecksFromApiV3(t, languageCode, clientID)

				assert.Equal(t, 1, len(result)) //hidden deck is absent
				for i := range result {
					checkDeckDTOFields(t, result[i])
				}
			})

			t.Run("get available decks after unlock", func(t *testing.T) {
				defer failOnPanic(t)

				promo := "pro"
				hiddenDeckID := "d4"
				languageCode := "RU"
				result := getDecksFromApiV3(t, languageCode, clientID)
				assert.Equal(t, 1, len(result)) //hidden deck is absent

				deck, status := enterPromo(t, promo, clientID)
				assert.NotNil(t, deck)
				assert.Equal(t, status, http.StatusCreated)
				checkDeckFields(t, deck)

				result = getDecksFromApiV3(t, languageCode, clientID)
				assert.Equal(t, 2, len(result))
				assert.Equal(t, hiddenDeckID, result[1].ID)
			})

			t.Run("get levels", func(t *testing.T) {
				defer failOnPanic(t)
				expected := []model.Level{
					expectedD1L1,
					expectedD1L2,
					expectedD1L3,
				}

				decks := getDecksFromApiV3(t, "EN", clientID)

				for i, deck := range decks {
					result := getLevelsFromApi(t, deck.ID, appUrl+apiV1)
					for _, level := range result {
						assert.NotEmpty(t, level.ID)
						assert.NotEmpty(t, level.DeckID)
						assert.NotEmpty(t, level.Name)
						assert.NotEmpty(t, level.LevelOrder)
					}
					if i == 0 {
						assert.Equal(t, expected, result)
					}
				}
			})

			t.Run("/deck/:id/levels - get levels", func(t *testing.T) {
				defer failOnPanic(t)
				clearHistory(t)
				expected := []output.LevelDto{
					{
						Level: expectedD1L1,
						Counts: &output.QuestionsCounts{
							QuestionsCount:       3,
							OpenedQuestionsCount: new(int),
						},
					},
					{
						Level: expectedD1L2,
						Counts: &output.QuestionsCounts{
							QuestionsCount:       2,
							OpenedQuestionsCount: new(int),
						},
					},
					{
						Level: expectedD1L3,
						Counts: &output.QuestionsCounts{
							QuestionsCount:       3,
							OpenedQuestionsCount: new(int),
						},
					},
				}

				result := getLevelsByDeckWithCountsFromApi(t, appUrl+apiV1, "d1", clientID)

				assert.NotNil(t, result)
				for i, level := range result {
					assert.Equal(t, expected[i].Level, level.Level)
					assert.Equal(t, expected[i].Counts.QuestionsCount, level.Counts.QuestionsCount)
					assert.Equal(t,
						*expected[i].Counts.OpenedQuestionsCount,
						*level.Counts.OpenedQuestionsCount)
				}
			})

			t.Run("/deck/:id/levels - check counts", func(t *testing.T) {
				defer failOnPanic(t)
				clearHistory(t)
				expected := []int{2, 2, 0}

				getQuestionFromApi(t, "d1l1", clientID, appUrl+apiV1)
				getQuestionFromApi(t, "d1l1", clientID, appUrl+apiV1)
				getQuestionFromApi(t, "d1l2", clientID, appUrl+apiV1)
				getQuestionFromApi(t, "d1l2", clientID, appUrl+apiV1)
				getQuestionFromApi(t, "d1l2", clientID, appUrl+apiV1) //not counted - more than in level

				result := getLevelsByDeckWithCountsFromApi(t, appUrl+apiV1, "d1", clientID)

				assert.NotNil(t, result)
				assert.NotEmpty(t, result)
				for i, level := range result {
					assert.NotNil(t, level)
					assert.Equal(t, expected[i], *level.Counts.OpenedQuestionsCount)
				}
			})

			t.Run("get level", func(t *testing.T) {
				defer failOnPanic(t)
				clearHistory(t)
				z := 0
				em11 := "em1"
				expected := output.LevelDto{
					Level: model.Level{
						ID:          "d1l1",
						DeckID:      "d1",
						LevelOrder:  1,
						Name:        "l1",
						Emoji:       &em11,
						ColorStart:  "0,0,0",
						ColorEnd:    "255,255,255",
						ColorButton: "1,1,1",
					},
					Counts: &output.QuestionsCounts{
						QuestionsCount:       3,
						OpenedQuestionsCount: &z,
					},
				}
				result := getLevelFromApi(t, appUrl+apiV1, "d1l1")
				assert.NotNil(t, result)
				assert.Equal(t, expected.Level, result.Level)
				assert.Equal(t, expected.Counts.QuestionsCount, result.Counts.QuestionsCount)
				assert.Equal(t,
					*expected.Counts.OpenedQuestionsCount,
					*result.Counts.OpenedQuestionsCount)
			})

			t.Run("get question", func(t *testing.T) {
				defer failOnPanic(t)
				question := getQuestionFromApi(t, "d1l1", clientID, appUrl+apiV1)
				assert.Contains(t, []string{"question d1l1q1 text", "question d1l1q2 text", "question d1l1q3 text"}, question.Text)
				assert.NotNil(t, question.ID)
				assert.NotNil(t, question.Text)
				assert.NotNil(t, question.LevelID)
				assert.NotNil(t, question.AdditionalText)
			})

			t.Run("questions in level are ordered (for many clients)", func(t *testing.T) {
				defer failOnPanic(t)
				clearHistory(t)
				questions := []string{"question d1l1q1 text", "question d1l1q2 text", "question d1l1q3 text"}

				wg := sync.WaitGroup{}

				for clientNum := 1; clientNum <= 5; clientNum++ {
					wg.Add(1)
					clientNumStr := strconv.Itoa(clientNum)
					go func() {
						defer wg.Done()
						for i := 0; i < 5; i++ {
							question := getQuestionFromApi(t, "d1l1", clientID+clientNumStr, appUrl+apiV1)

							ansIndex1 := utils.FindIndex(questions, question.Text)
							assert.NotEqual(t, -1, ansIndex1)

							question = getQuestionFromApi(t, "d1l1", clientID+clientNumStr, appUrl+apiV1)
							ansIndex2 := utils.FindIndex(questions, question.Text)
							assert.NotEqual(t, -1, ansIndex2)
							assert.NotEqual(t, ansIndex1, ansIndex2)

							question = getQuestionFromApi(t, "d1l1", clientID+clientNumStr, appUrl+apiV1)
							ansIndex3 := utils.FindIndex(questions, question.Text)
							assert.NotEqual(t, -1, ansIndex3)
							assert.NotEqual(t, ansIndex1, ansIndex3)
							assert.NotEqual(t, ansIndex2, ansIndex3)
							println("ORDER CHECK FINISHED")
							time.Sleep(100 * time.Millisecond)
						}
					}()
				}
				wg.Wait()
			})
			t.Run("Get all questions by deck", func(t *testing.T) {
				defer failOnPanic(t)
				clearHistory(t)
				questions := getAllQuestionsFromDeck(t, "d1", appUrl+apiV1)
				assert.NotNil(t, questions)
				assert.Equal(t, 8, len(questions))
				for _, q := range questions {
					assert.True(t, strings.HasPrefix(q.Text, "question d1"))
					assert.NotNil(t, q.ID)
					assert.NotNil(t, q.LevelID)
					assert.NotNil(t, q.AdditionalText)
				}
			})

			t.Run("Get vector images", func(t *testing.T) {
				defer failOnPanic(t)
				clearHistory(t)
				imageContent := getVectorImage(t, "1", appUrl+apiV1)
				assert.Equal(t, "<svg>1</svg>", imageContent)
				imageContent = getVectorImage(t, "2", appUrl+apiV1)
				assert.Equal(t, "<svg>2</svg>", imageContent)
			})
		})
	}
}

func getVectorImage(t *testing.T, id string, url string) string {
	fmt.Println("Getting vector image " + id)
	request, err := http.NewRequest("GET", url+"/get-vector-image/"+id, nil)
	assert.NoError(t, err)

	client := http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result string
	all, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	result = string(all)
	err = response.Body.Close()
	return result
}

func getAllQuestionsFromDeck(t *testing.T, deckID string, url string) []model.Question {
	fmt.Println("Getting levels of deck " + deckID)
	request, err := http.NewRequest("GET", url+"/deck/"+deckID+"/questions", nil)
	assert.NoError(t, err)

	client := http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result []model.Question
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()
	return result
}

func getQuestionFromApi(t *testing.T, levelID string, clientID string, url string) *model.Question {
	fmt.Println("Getting question from level " + levelID)
	request, err := http.NewRequest("GET", url+"/question", nil)
	assert.NoError(t, err)

	query := request.URL.Query()
	query.Add("levelId", levelID)
	query.Add("clientId", clientID)
	request.URL.RawQuery = query.Encode()
	client := http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result model.Question
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()
	println("Got question " + result.Text + " for client " + clientID)
	return &result
}

func getLevelFromApi(t *testing.T, url, id string) output.LevelDto {
	fmt.Println("Getting level with id " + id)
	request, err := http.NewRequest("GET", url+"/level/"+id, nil)
	assert.NoError(t, err)

	query := request.URL.Query()
	query.Add("clientId", clientID)
	query.Add("features", "OPENED_COUNT")
	request.URL.RawQuery = query.Encode()
	client := http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result output.LevelDto
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()
	return result
}

func getLevelsFromApi(t *testing.T, deckID string, url string) []model.Level {
	fmt.Println("Getting levels of deck " + deckID)
	request, err := http.NewRequest("GET", url+"/levels", nil)
	assert.NoError(t, err)

	query := request.URL.Query()
	query.Add("deckId", deckID)
	request.URL.RawQuery = query.Encode()
	client := http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result []model.Level
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()
	return result
}

func getLevelsByDeckWithCountsFromApi(t *testing.T, url, deckID, userID string) []output.LevelDto {
	fmt.Println("Getting levels of deck " + deckID + " for user " + userID)
	request, err := http.NewRequest("GET", url+"/deck/"+deckID+"/levels", nil)
	assert.NoError(t, err)

	query := request.URL.Query()
	query.Add("clientId", clientID)
	request.URL.RawQuery = query.Encode()
	client := http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result []output.LevelDto
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()
	return result
}

/*func getDecksFromApi(t *testing.T, url string, lang string) []output.DeckDTO {
	fmt.Println("Getting decks...")
	request, err := http.NewRequest("GET", url+"/decks?languageCode="+lang, nil)
	assert.NoError(t, err)

	client := http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result []output.DeckDTO
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()

	return result
}*/

func getDecksFromApiV3(t *testing.T, languageCode string, clientId string) []output.DeckDTO {
	fmt.Println("Getting localized decks...")
	url := config.GetConfig("TEST_URL")
	request, err := http.NewRequest("GET", url+"/api/v3/decks?languageCode="+languageCode+"&clientId="+clientId, nil)
	assert.NoError(t, err)

	client := http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result []output.DeckDTO
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()

	return result
}

func enterPromo(t *testing.T, promo, clientId string) (model.Deck, int) {
	fmt.Println("Getting localized decks...")
	url := getTestUrl()
	request, err := http.NewRequest("POST", url+"/api/v1/enter-promo/"+promo+"?clientId="+clientId, nil)
	assert.NoError(t, err)

	client := http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Contains(t, []int{http.StatusCreated, http.StatusNoContent}, response.StatusCode)

	var result model.Deck
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()

	return result, response.StatusCode
}

func getTestUrl() string {
	return config.GetConfig("TEST_URL")
}

func getResponseCode(method, url string) (int, error) {
	fmt.Printf("Checking status from request %s\n", url)
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 0, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil
}

func checkDeckFields(t *testing.T, deck model.Deck) {
	assert.NotNil(t, deck.ID)
	assert.NotNil(t, deck.LanguageCode)
	assert.NotNil(t, deck.Name)
	assert.NotNil(t, deck.Description)
	assert.NotNil(t, deck.Labels)
}
func checkDeckDTOFields(t *testing.T, deck output.DeckDTO) {
	checkDeckFields(t, deck.Deck)
	assert.NotNil(t, deck.CardsCount)
	assert.NotNil(t, deck.OpenedCount)
}

func failOnPanic(t *testing.T) {
	if r := recover(); r != nil {
		t.Fatalf("The code panicked: %v", r)
	}
}

func createUpdateObject(text string) *tgbotapi.Update {
	firstName := "German"
	lastName := "Reus"
	userName := "Logotipiwe"
	user := &tgbotapi.User{
		ID:           int64(rand.Int()),
		IsBot:        false,
		FirstName:    firstName,
		LastName:     lastName,
		UserName:     userName,
		LanguageCode: "en",
	}

	chat := &tgbotapi.Chat{
		ID:        1111111,
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
		Type:      "private",
	}

	currentTime := int(time.Now().Unix())
	message := &tgbotapi.Message{
		MessageID: rand.Int(),
		From:      user,
		Chat:      chat,
		Date:      currentTime,
		Text:      text,
	}

	update := &tgbotapi.Update{
		UpdateID: rand.Int(),
		Message:  message,
	}
	return update
}

func createUpdateObjectFrom(update *tgbotapi.Update, text string) *tgbotapi.Update {
	currentTime := int(time.Now().Unix())

	update.Message.Text = text
	update.Message.Date = currentTime
	update.Message.MessageID = rand.Int()
	return update
}

func toMarkup(t *testing.T, input interface{}) *tgbotapi.ReplyKeyboardMarkup {
	var ans tgbotapi.ReplyKeyboardMarkup
	jsonn, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(jsonn, &ans)
	if err != nil {
		t.Fatal(err)
	}
	return &ans
}

func sendUpdate(t *testing.T, update *tgbotapi.Update) *tgbotapi.MessageConfig {
	appUrl := config.GetConfig("TEST_URL")
	println("Url to test " + appUrl)
	m2mToken := config.GetConfig("M_TOKEN")

	println("sending message " + update.Message.Text)
	reqBody, err := json.Marshal(update)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", appUrl+http2.IntegrationTestPrefix+"/test-chat", bytes.NewReader(reqBody))
	assert.NoError(t, err)

	query := req.URL.Query()
	query.Add("mToken", m2mToken)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	response, err := client.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result tgbotapi.MessageConfig
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()
	if err != nil {
		panic(err)
	}

	println("Got and decoded answer")

	return &result
}

func clearHistory(t *testing.T) {
	appUrl := config.GetConfig("TEST_URL")
	println("Clearing questions history")
	req, err := http.NewRequest("POST", appUrl+http2.IntegrationTestPrefix+"/clear-history", nil)

	m2mToken := config.GetConfig("M_TOKEN")
	query := req.URL.Query()
	query.Add("mToken", m2mToken)
	req.URL.RawQuery = query.Encode()

	assert.NoError(t, err)
	client := &http.Client{}
	response, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	println("History cleared")
}

func checkIfImagesEnabled(t *testing.T) {
	appUrl := config.GetConfig("TEST_URL")
	println("Checking if images are enabled...")

	req, err := http.NewRequest("GET", appUrl+http2.IntegrationTestPrefix+"/images-enabled", nil)
	assert.NoError(t, err)

	m2mToken := config.GetConfig("M_TOKEN")
	query := req.URL.Query()
	query.Add("mToken", m2mToken)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	response, err := client.Do(req)
	assert.NoError(t, err, "FUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUU")

	assert.Equal(t, http.StatusOK, response.StatusCode)

	bodyBytes, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	result := string(bodyBytes)
	resultBool, err := strconv.ParseBool(result)
	assert.NoError(t, err)
	err = response.Body.Close()
	assert.NoError(t, err)

	if resultBool {
		err = errors.New("error: images enabled. Cannot perform integration tests")
		t.Fatal(err)
	} else {
		println("Images are disabled, start testing...")
	}
}
