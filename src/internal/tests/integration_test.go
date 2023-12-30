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

func TestApplication(t *testing.T) {
	res, err := strconv.ParseBool(config.GetConfigOr("DO_INTEGRATION_TESTS", "true"))
	if err == nil && !res {
		//couldn't find way to not execute it in unit tests run
		println("TEST SKIPPED")
		return
	}

	checkIfImagesEnabled(t)

	t.Run("Telegram client", func(t *testing.T) {
		t.Run("start message", func(t *testing.T) {
			defer failOnPanic(t)
			update := createUpdateObject("/start")
			ans := sendUpdate(t, update)
			assert.Equal(t, ans.Text, service.WelcomeMessage)
		})

		t.Run("get decks start", func(t *testing.T) {
			defer failOnPanic(t)
			update := createUpdateObject("/start")
			ans := sendUpdate(t, update)
			replyMarkup := toMarkup(t, ans.BaseChat.ReplyMarkup)
			assert.Equal(t, 3, len(replyMarkup.Keyboard))
			assert.Equal(t, "deck d1 name", replyMarkup.Keyboard[0][0].Text)
			assert.Equal(t, "deck d2 name", replyMarkup.Keyboard[1][0].Text)
			assert.Equal(t, "deck d3 name", replyMarkup.Keyboard[2][0].Text)
			println(ans)
		})

		t.Run("select deck", func(t *testing.T) {
			defer failOnPanic(t)
			update := createUpdateObject("/start")
			ans := sendUpdate(t, update)
			update = createUpdateObject("deck d1 name")
			ans = sendUpdate(t, update)
			replyMarkup := toMarkup(t, ans.BaseChat.ReplyMarkup)
			assert.Equal(t, ans.Text, service.GotLevelsMessage)
			assert.Equal(t, 3, len(replyMarkup.Keyboard[0]))
			assert.Equal(t, "l1", replyMarkup.Keyboard[0][0].Text)
			assert.Equal(t, "l2", replyMarkup.Keyboard[0][1].Text)
			assert.Equal(t, "l3", replyMarkup.Keyboard[0][2].Text)
			println(ans)
		})

		t.Run("select deck; select level", func(t *testing.T) {
			defer failOnPanic(t)
			update := createUpdateObject("/start")
			ans := sendUpdate(t, update)
			update = createUpdateObjectFrom(update, "deck d1 name")
			ans = sendUpdate(t, update)
			update = createUpdateObjectFrom(update, "l1")
			ans = sendUpdate(t, update)
			assert.Contains(t, []string{"question d1l1q1 text", "question d1l1q2 text", "question d1l1q3 text"}, ans.Text)
			println(ans)
		})

		t.Run("select level > markup nil", func(t *testing.T) {
			defer failOnPanic(t)
			update := createUpdateObject("/start")
			ans := sendUpdate(t, update)
			update = createUpdateObjectFrom(update, "deck d1 name")
			ans = sendUpdate(t, update)
			update = createUpdateObjectFrom(update, "l1")
			ans = sendUpdate(t, update)
			assert.Nil(t, ans.BaseChat.ReplyMarkup)
			println(ans)
		})

		t.Run("select deck; select level many times", func(t *testing.T) {
			defer failOnPanic(t)
			update := createUpdateObject("/start")
			ans := sendUpdate(t, update)
			update = createUpdateObjectFrom(update, "deck d1 name")
			ans = sendUpdate(t, update)
			for i := 0; i < 10; i++ {
				update = createUpdateObjectFrom(update, "l1")
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
				update = createUpdateObjectFrom(update, "l1")
				ans = sendUpdate(t, update)
				ansIndex1 := utils.FindIndex(questions, ans.Text)
				assert.NotEqual(t, -1, ansIndex1)

				update = createUpdateObjectFrom(update, "l1")
				ans = sendUpdate(t, update)
				ansIndex2 := utils.FindIndex(questions, ans.Text)
				assert.NotEqual(t, -1, ansIndex2)
				assert.NotEqual(t, ansIndex1, ansIndex2)

				update = createUpdateObjectFrom(update, "l1")
				ans = sendUpdate(t, update)
				ansIndex3 := utils.FindIndex(questions, ans.Text)
				assert.NotEqual(t, -1, ansIndex3)
				assert.NotEqual(t, ansIndex1, ansIndex3)
				assert.NotEqual(t, ansIndex2, ansIndex3)
				println("ORDER CHECK FINISHED")
				time.Sleep(100 * time.Millisecond)
			}
		})

		t.Run("select deck; select different levels many times", func(t *testing.T) {
			defer failOnPanic(t)
			update := createUpdateObject("/start")
			ans := sendUpdate(t, update)
			update = createUpdateObjectFrom(update, "deck d1 name")
			ans = sendUpdate(t, update)
			for i := 0; i < 20; i++ {
				level := strconv.Itoa(rand.Intn(3) + 1)
				update = createUpdateObjectFrom(update, "l"+level)
				ans = sendUpdate(t, update)
				assert.True(t, strings.HasPrefix(ans.Text, "question d1l"+level))
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

		t.Run("/feedback command", func(t *testing.T) {
			defer failOnPanic(t)
			update := createUpdateObject(service.FeedbackCommand)
			ans := sendUpdate(t, update)
			assert.Equal(t, ans.Text, service.AssignFeedbackText)
			update = createUpdateObjectFrom(update, "MyFeedback")
			ans = sendUpdate(t, update)
			assert.Equal(t, ans.Text, service.AcceptFeedbackText)
		})
	})

	t.Run("Api client", func(t *testing.T) {
		appUrl := config.GetConfig("TEST_URL")
		println("Url to test " + appUrl)
		apiV1 := "/api/v1"

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

			expected := [][]string{
				{"l1", "l2", "l3"},
				{"l1", "l2", "l3"},
				{"l1", "l2"},
			}

			decks := getDecksFromApi(t, appUrl+apiV1)

			for i, deck := range decks {
				result := getLevelsFromApi(t, deck.ID, appUrl+apiV1)

				assert.Equal(t, expected[i], result)
			}
		})

		t.Run("get question", func(t *testing.T) {
			defer failOnPanic(t)
			question := getQuestionFromApi(t, "d1", "l1", appUrl+apiV1)
			assert.Contains(t, []string{"question d1l1q1 text", "question d1l1q2 text", "question d1l1q3 text"}, question.Text)
			assert.NotNil(t, question.ID)
			assert.NotNil(t, question.Text)
			assert.NotNil(t, question.DeckID)
			assert.NotNil(t, question.Level)
		})

		t.Run("questions in level are ordered", func(t *testing.T) {
			defer failOnPanic(t)
			clearHistory(t)
			questions := []string{"question d1l1q1 text", "question d1l1q2 text", "question d1l1q3 text"}
			for i := 0; i < 5; i++ {
				question := getQuestionFromApi(t, "d1", "l1", appUrl+apiV1)

				ansIndex1 := utils.FindIndex(questions, question.Text)
				assert.NotEqual(t, -1, ansIndex1)

				question = getQuestionFromApi(t, "d1", "l1", appUrl+apiV1)
				ansIndex2 := utils.FindIndex(questions, question.Text)
				assert.NotEqual(t, -1, ansIndex2)
				assert.NotEqual(t, ansIndex1, ansIndex2)

				question = getQuestionFromApi(t, "d1", "l1", appUrl+apiV1)
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

func getQuestionFromApi(t *testing.T, deckID string, levelName string, url string) *domain.Question {
	fmt.Println("Getting question from deck " + deckID + ", level " + levelName)
	request, err := http.NewRequest("GET", url+"/question", nil)
	assert.NoError(t, err)

	query := request.URL.Query()
	query.Add("deckId", deckID)
	query.Add("levelName", levelName)
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

func getLevelsFromApi(t *testing.T, deckID string, url string) []string {
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

	var result []string
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
	assert.NoError(t, err)

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
