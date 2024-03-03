package tests

import (
	"bytes"
	http2 "dc_haur/src/http"
	"dc_haur/src/internal/domain"
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
	"testing"
	"time"
)

const d1l1QuestionID = "4f84bde5-d6ad-4a2d-a2da-0553b4b281a2"

const (
	d1   = "em1 deck d1 name"
	d2   = "em2 deck d2 name"
	d3   = "deck d3 name"
	d1l1 = "em1 l1"
	d1l2 = "em1 l2"
	d1l3 = "l3"
	d2l1 = "em2 l1"
	d2l2 = "l2"
	d2l3 = "em2 l3"
	d3l1 = "em3 l1"
	d3l2 = "l2"
)

func TestApplication(t *testing.T) {

	res, err := strconv.ParseBool(config.GetConfigOr("DO_INTEGRATION_TESTS", "true"))
	if err == nil && !res {
		//couldn't find way to not execute it in unit tests run
		println("TEST SKIPPED")
		return
	}

	checkIfImagesEnabled(t)

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
				assert.Equal(t, 3, len(replyMarkup.Keyboard))
				assert.Equal(t, d1, replyMarkup.Keyboard[0][0].Text)
				assert.Equal(t, d2, replyMarkup.Keyboard[1][0].Text)
				assert.Equal(t, d3, replyMarkup.Keyboard[2][0].Text)
				println(ans)
			})

			t.Run("get decks start from group", func(t *testing.T) {
				defer failOnPanic(t)
				update := createUpdateObject("/start@HowAreYouReallyGameBot")
				ans := sendUpdate(t, update)
				replyMarkup := toMarkup(t, ans.BaseChat.ReplyMarkup)
				assert.Equal(t, 3, len(replyMarkup.Keyboard))
				assert.Equal(t, d1, replyMarkup.Keyboard[0][0].Text)
				assert.Equal(t, d2, replyMarkup.Keyboard[1][0].Text)
				assert.Equal(t, d3, replyMarkup.Keyboard[2][0].Text)
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

			t.Run("get decks", func(t *testing.T) {
				defer failOnPanic(t)

				result := getDecksFromApi(t, appUrl+apiV1)

				assert.Equal(t, 3, len(result))
				for i := range result {
					assert.NotNil(t, result[i].ID)
					assert.NotNil(t, result[i].Name)
					assert.NotNil(t, result[i].Description)
				}
			})

			t.Run("get levels", func(t *testing.T) {
				defer failOnPanic(t)
				em11 := "em1"
				expected := []domain.Level{
					{
						ID:         "4f84bde5-d6ad-4a2d-a2da-0553b4b281a2",
						DeckID:     "d1",
						LevelOrder: 1,
						Name:       "l1",
						Emoji:      &em11,
						ColorStart: "0,0,0",
						ColorEnd:   "255,255,255",
					},
					{
						ID:         "dae6f634-8a6c-42a7-8d25-6a44e91e6e21",
						DeckID:     "d1",
						LevelOrder: 2,
						Name:       "l2",
						Emoji:      &em11,
						ColorStart: "0,0,0",
						ColorEnd:   "255,255,255",
					},
					{
						ID:         "8e7e1f07-0292-4ef6-8529-fb92a0d4c1f6",
						DeckID:     "d1",
						LevelOrder: 3,
						Name:       "l3",
						Emoji:      nil,
						ColorStart: "0,0,0",
						ColorEnd:   "255,255,255",
					},
				}

				decks := getDecksFromApi(t, appUrl+apiV1)

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

			t.Run("get question", func(t *testing.T) {
				defer failOnPanic(t)
				question := getQuestionFromApi(t, d1l1QuestionID, appUrl+apiV1)
				assert.Contains(t, []string{"question d1l1q1 text", "question d1l1q2 text", "question d1l1q3 text"}, question.Text)
				assert.NotNil(t, question.ID)
				assert.NotNil(t, question.Text)
				assert.NotNil(t, question.LevelID)
			})

			t.Run("questions in level are ordered", func(t *testing.T) {
				defer failOnPanic(t)
				clearHistory(t)
				questions := []string{"question d1l1q1 text", "question d1l1q2 text", "question d1l1q3 text"}
				for i := 0; i < 5; i++ {
					question := getQuestionFromApi(t, d1l1QuestionID, appUrl+apiV1)

					ansIndex1 := utils.FindIndex(questions, question.Text)
					assert.NotEqual(t, -1, ansIndex1)

					question = getQuestionFromApi(t, d1l1QuestionID, appUrl+apiV1)
					ansIndex2 := utils.FindIndex(questions, question.Text)
					assert.NotEqual(t, -1, ansIndex2)
					assert.NotEqual(t, ansIndex1, ansIndex2)

					question = getQuestionFromApi(t, d1l1QuestionID, appUrl+apiV1)
					ansIndex3 := utils.FindIndex(questions, question.Text)
					assert.NotEqual(t, -1, ansIndex3)
					assert.NotEqual(t, ansIndex1, ansIndex3)
					assert.NotEqual(t, ansIndex2, ansIndex3)
					println("ORDER CHECK FINISHED")
					time.Sleep(100 * time.Millisecond)
				}
			})
		})
	}
}

func getQuestionFromApi(t *testing.T, levelID string, url string) *domain.Question {
	fmt.Println("Getting question from level " + levelID)
	request, err := http.NewRequest("GET", url+"/question", nil)
	assert.NoError(t, err)

	query := request.URL.Query()
	query.Add("levelId", levelID)
	query.Add("clientId", "integrationTestsClient")
	request.URL.RawQuery = query.Encode()
	client := http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result domain.Question
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()
	return &result
}

func getLevelsFromApi(t *testing.T, deckID string, url string) []domain.Level {
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

	var result []domain.Level
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()
	return result
}

func getDecksFromApi(t *testing.T, url string) []domain.Deck {
	fmt.Println("Getting decks...")
	request, err := http.NewRequest("GET", url+"/decks", nil)
	assert.NoError(t, err)

	client := http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	var result []domain.Deck
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.NoError(t, err)
	err = response.Body.Close()

	return result
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
