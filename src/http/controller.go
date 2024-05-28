package http

import (
	_ "dc_haur/docs"
	"dc_haur/src/internal/model"
	"dc_haur/src/internal/model/output"
	"dc_haur/src/internal/repo"
	"dc_haur/src/internal/service"
	"dc_haur/src/pkg"
	"errors"
	"github.com/Logotipiwe/dc_go_auth_lib/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	config "github.com/logotipiwe/dc_go_config_lib"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"image/png"
	"log"
	"net/http"
	"strings"
)

const IntegrationTestPrefix = "/api/v1/integration-test"

type Controller struct {
	services *service.Services
}

func StartServer(services *service.Services) {
	controller := Controller{services: services}
	router := gin.Default()
	integrationTestingRoutes := router.Group(IntegrationTestPrefix)

	integrationTestingRoutes.Use(func(c *gin.Context) {
		if err := auth.AuthAsMachine(c.Request); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	})

	integrationTestingRoutes.GET("/test-image", doWithErrExplicit(func(c *gin.Context) error {
		additionalText := "Отвечает человек слева"
		question := model.Question{
			ID:             "1",
			LevelID:        "2",
			Text:           "Как ты думаешь, что самое сложное в том деле, которым я зарабатываю себе на жизнь?",
			AdditionalText: &additionalText,
		}
		card, err := service.CreateImageCardFromQuestion(&question, "", "")
		if err != nil {
			return err
		}
		c.Writer.Header().Set("Content-Type", "image/png")
		err = png.Encode(c.Writer, card)
		if err != nil {
			return err
		}
		return nil
	}))
	integrationTestingRoutes.POST("/test-chat", doWithErrExplicit(func(c *gin.Context) error {
		var update tgbotapi.Update
		if err := c.ShouldBindJSON(&update); err != nil {
			return err
		}
		reply, err := services.TgUpdatesHandler.HandleMessageAndReply(update)
		if err != nil {
			reply = services.TgUpdatesHandler.SendUnknownCommandAnswer(update)
		}
		c.JSON(http.StatusOK, reply)
		return nil
	}))
	integrationTestingRoutes.POST("/clear-history", doWithErrExplicit(func(c *gin.Context) error {
		if err := services.Repos.History.Truncate(); err != nil {
			return err
		}
		if err := services.Repos.UsedQuestions.Truncate(); err != nil {
			return err
		}
		if err := services.Repos.Decks.TruncateUnlockedDecks(); err != nil {
			return err
		}
		if err := services.Repos.QuestionReactions.Truncate(); err != nil {
			return err
		}
		if err := services.Repos.DeckLikes.Truncate(); err != nil {
			return err
		}
		c.Status(http.StatusOK)
		return nil
	}))
	integrationTestingRoutes.GET("/images-enabled", doWithErrExplicit(func(c *gin.Context) error {
		enabledImagesStr := config.GetConfig("ENABLE_IMAGES")
		if _, err := c.Writer.WriteString(enabledImagesStr); err != nil {
			return err
		}
		return nil
	}))

	apiV1 := router.Group("/api/v1")
	apiV2 := router.Group("/api/v2")
	apiV3 := router.Group("/api/v3")

	apiV1.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	})
	apiV2.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	})
	apiV3.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1.GET("/levels", doWithErr(controller.GetLevels))
	apiV1.GET("deck/:deckId/levels", doWithErr(controller.GetLevelsWithCounts))
	apiV1.GET("/level/:id", doWithErr(controller.GetLevel))
	apiV1.GET("/question", doWithErr(controller.GetQuestion))
	apiV1.POST("/question/:questionId/like", doWithErr(controller.LikeQuestion))
	apiV1.POST("/question/:questionId/dislike", doWithErr(controller.DislikeQuestion))
	apiV1.GET("/deck/:deckId/questions", doWithErr(controller.GetDeckQuestions))
	apiV1.POST("/deck/:deckId/like", doWithErr(controller.LikeDeck))
	apiV1.POST("/deck/:deckId/dislike", doWithErr(controller.DislikeDeck))
	apiV1.GET("/get-vector-image/:id", doWithErr(controller.GetImage))
	apiV1.POST("/enter-promo/:promo", doWithErr(controller.EnterPromo))

	apiV1.GET("/user/:userId/likes", doWithErr(controller.GetUserLikesDEPRECATED))

	apiV2.GET("/decks", doWithErr(controller.GetLocalizedDecks))
	apiV2.POST("/question/:questionId/react-like", doWithErr(controller.LikeQuestion))
	apiV2.POST("/question/:questionId/react-dislike", doWithErr(controller.DislikeQuestionV2))
	apiV2.POST("/question/:questionId/react-remove", doWithErr(controller.RemoveReactionToQuestion))
	apiV2.GET("/user/:userId/reactions", doWithErr(controller.GetUserReactions))

	apiV3.GET("/decks", doWithErr(controller.GetLocalizedAvailableDecks))

	port := config.GetConfigOr("CONTAINER_PORT", "80")
	log.Println("Starting server on port " + port)
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

// GetLocalizedDecks godoc
// @Summary      Get decks by lang code
// @Param 		 languageCode query string true "Language code in upper case (RU, EN)"
// @Produce      json
// @Success      200  {array} output.DeckDTO
// @Router       /v2/decks [get]
// TODO deprecated
func (c Controller) GetLocalizedDecks(ctx *gin.Context) error {
	langCode := ctx.Query("languageCode")
	if langCode == "" {
		ctx.String(400, "Language not specified. "+
			"Please specify languageCode query parameter as in the following: languageCode=EN")
		return nil
	}

	decksService := c.services.Decks
	models, err := decksService.GetDecksByLanguage(langCode)
	if err != nil {
		return err
	}
	decks, err := decksService.EnrichDecksWithCardsCounts(decksService.ToDtos(models))
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, decks)
	return nil
}

// GetLocalizedAvailableDecks godoc
// @Summary      Get decks by lang code
// @Param 		 languageCode query string true "Language code in upper case (RU, EN)"
// @Param 		 clientId query string true "Language code in upper case (RU, EN)"
// @Produce      json
// @Success      200  {array} output.DeckDTO
// @Router       /v3/decks [get]
func (c Controller) GetLocalizedAvailableDecks(ctx *gin.Context) error {
	langCode := ctx.Query("languageCode")
	if langCode == "" {
		ctx.String(400, "Language not specified. "+
			"Please specify languageCode query parameter as in the following: languageCode=EN")
		return nil
	}
	clientID := ctx.Query("clientId")
	if clientID == "" {
		return errors.New("please specify clientId")
	}

	decksService := c.services.Decks
	models, err := decksService.GetAvailableForUserDecksByLanguage(langCode, clientID)
	if err != nil {
		return err
	}
	decks, err := decksService.EnrichDecksWithCounts(models, clientID)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, decks)
	return nil
}

// GetLevels godoc
// @Summary      ПЕРЕХОДИТЬ НА /deck/:id/levels. Получить уровни в колоде.
// @Param 		 deckId query string true "Id of deck for which selecting levels"
// @Produce      json
// @Success      200  {array} model.Level
// @Router       /v1/levels [get]
func (c Controller) GetLevels(ctx *gin.Context) error {
	deckId := ctx.Query("deckId")
	levels, err := c.services.Repos.Levels.GetLevelsByDeckId(deckId)
	if errors.Is(err, repo.NoLevelsErr) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No levels found by deck id " + deckId})
		return nil
	} else if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, levels)
	return nil
}

// GetLevelsWithCounts godoc
// @Summary      Уровни в колоде. С кличеством просмотренных карт.
// @Param 		 deckId path string true "Id of deck for which selecting levels"
// @Param 		 clientId query string true "client id. Нужен чтобы получить количество открытых вопросов"
// @Produce      json
// @Success      200  {array} model.Level
// @Router       /v1/deck/{deckId}/levels [get]
func (c Controller) GetLevelsWithCounts(ctx *gin.Context) error {
	deckId := ctx.Param("deckId")
	clientId := ctx.Query("clientId")
	dtos, err := c.services.Levels.GetLevelsOfDeckWithCounts(deckId, clientId)
	if errors.Is(err, repo.NoLevelsErr) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No levels found by deck id " + deckId})
		return nil
	} else if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, dtos)
	return nil
}

// GetLevel godoc
// @Summary      НЕ ИСПОЛЬЗУЕТСЯ. Получить уровень по id (с возможностью получить кол-во просмотренных карт)
// @Param 		 clientId query string false "id клиента. Обязателен если в features есть OPENED_COUNT"
// @Param 		 features query string false "Список модификаций ответа через запятую. Доступны - OPENED_COUNT"
// @Param 		 id path string true "id of level"
// @Produce      json
// @Success      200  {object} output.LevelDto
// @Router       /v1/level/{id} [get]
func (c Controller) GetLevel(ctx *gin.Context) error {
	levelID := ctx.Param("id")
	level, err := c.services.Repos.Levels.GetByID(levelID)
	if err != nil {
		return err
	}

	dto := output.LevelDto{Level: level}
	dto, err = c.services.Levels.EnrichWithCount(dto)
	if err != nil {
		return err
	}

	features := strings.Split(ctx.Query("features"), ",")
	if pkg.Includes("OPENED_COUNT", features) {
		clientID := ctx.Query("clientId")
		if clientID == "" {
			return errors.New("client id is empty")
		}
		dto, err = c.services.Levels.EnrichWithOpenedCount(dto, clientID)
		if err != nil {
			return err
		}
	}
	ctx.JSON(http.StatusOK, dto)
	return nil
}

// GetQuestion godoc
// @Summary      Get random question by selected level
// @Param 		 levelId query string true "Id of level for which selecting question"
// @Param 		 clientId query string true "Client id - differs clients from each other. Needed for ordering random questions for each client/"
// @Produce      json
// @Success      200  {object} model.Question
// @Router       /v1/question [get]
func (c Controller) GetQuestion(ctx *gin.Context) error {
	levelID := ctx.Query("levelId")
	clientId, _ := ctx.GetQuery("clientId")
	if clientId == "" {
		return errors.New("you must specify clientId")
	}
	question, isLast, err := c.services.Questions.GetRandQuestion(levelID, clientId)
	if err != nil {
		return err
	}
	dto := output.QuestionDTO{
		Question: *question,
		IsLast:   isLast,
	}
	ctx.JSON(http.StatusOK, dto)
	return nil
}

// GetDeckQuestions godoc
// @Summary      Get all questions from specified deck
// @Quer
// @Param 		 deckId path string true "Id of deck for which questions are selected"
// @Produce      json
// @Success      200  {array} model.Question
// @Router       /v1/deck/{deckId}/questions [get]
func (c Controller) GetDeckQuestions(ctx *gin.Context) error {
	deckId := ctx.Param("deckId")
	questions, err := c.services.Repos.Questions.GetAllByDeckId(deckId)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, questions)
	return nil
}

// GetImage godoc
// @Summary      Get code of vector image
// @Param 		 id path string true "Id of image"
// @Produce      xml
// @Success      200 {xml} svg
// @Router       /v1/get-vector-image/{id} [get]
func (c Controller) GetImage(ctx *gin.Context) error {
	imageId := ctx.Param("id")
	image, err := c.services.Repos.VectorImages.GetVectorImageById(imageId)
	if err != nil {
		return err
	}
	ctx.Header("Content-Type", "image/svg+xml")
	ctx.String(200, image.Content)
	if err != nil {
		return err
	}
	return nil
}

// LikeQuestion godoc
// @Summary      Like a question
// @Description  Endpoint to like a particular question. Gives 409 in case of duplicating like
// @Produce      json
// @Param        questionId path string true "Question ID"
// @Param        userId query string true "User ID"
// @Success      200 {object} map[string]string
// @Failure      400,409 {object} map[string]string
// @Router       /v1/question/{questionId}/like [post]
// @Router       /v2/question/{questionId}/react-like [post]
func (c Controller) LikeQuestion(ctx *gin.Context) error {
	questionId := ctx.Param("questionId")
	userId := ctx.Query("userId")

	if _, err := uuid.Parse(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID format for userId",
		})
		return nil
	}

	err := c.services.QuestionReactionsService.Like(userId, questionId)
	if err != nil {
		var driverErr *mysql.MySQLError
		if errors.As(err, &driverErr) {
			if driverErr.Number == pkg.SqlDuplicateErrState {
				ctx.JSON(http.StatusConflict, gin.H{"error": "Like duplicated"})
				return nil
			}
		}

		return err
	}
	ctx.Status(http.StatusOK)
	return nil

}

// DislikeQuestion godoc
// @Summary      DislikeDEPRECATED a question
// @Description  Endpoint remove like from a particular question
// @Produce      json
// @Param        questionId path string true "Question ID"
// @Param        userId query string true "User ID"
// @Success      200 {object} map[string]string
// @Failure      400,500 {object} map[string]string
// @Router       /v1/question/{questionId}/dislike [post]
func (c Controller) DislikeQuestion(ctx *gin.Context) error {
	questionId := ctx.Param("questionId")
	userId := ctx.Query("userId")

	if _, err := uuid.Parse(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID format for userId",
		})
		return nil
	}

	err := c.services.QuestionReactionsService.DislikeDEPRECATED(userId, questionId)
	if err != nil {
		return err
	}
	ctx.Status(http.StatusOK)
	return nil
}

// RemoveReactionToQuestion godoc
// @Summary      Undo like a question
// @Description  Endpoint remove like from a particular question
// @Produce      json
// @Param        questionId path string true "Question ID"
// @Param        userId query string true "User ID"
// @Success      200 {object} map[string]string
// @Failure      400,500 {object} map[string]string
// @Router       /v2/question/{questionId}/react-remove [post]
func (c Controller) RemoveReactionToQuestion(ctx *gin.Context) error {
	questionId := ctx.Param("questionId")
	userId := ctx.Query("userId")

	if _, err := uuid.Parse(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID format for userId",
		})
		return nil
	}

	err := c.services.QuestionReactionsService.RemoveReaction(userId, questionId)
	if err != nil {
		return err
	}
	ctx.Status(http.StatusOK)
	return nil
}

// DislikeQuestionV2 godoc
// @Summary      Set dislike to a question
// @Description  Endpoint to dislike a specific question. Gives 409 in case of duplicating dislike
// @Produce      json
// @Param        questionId path string true "question ID"
// @Param        userId query string true "User ID"
// @Success      200 {object} map[string]string
// @Failure      400,409 {object} map[string]string
// @Router       /v2/question/{questionId}/react-dislike [post]
func (c Controller) DislikeQuestionV2(ctx *gin.Context) error {
	questionId := ctx.Param("questionId")
	userId := ctx.Query("userId")

	if _, err := uuid.Parse(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID format for userId",
		})
		return nil
	}

	err := c.services.QuestionReactionsService.DislikeV2(userId, questionId)

	if err != nil {
		var driverErr *mysql.MySQLError
		if errors.As(err, &driverErr) {
			if driverErr.Number == pkg.SqlDuplicateErrState {
				ctx.JSON(http.StatusConflict, gin.H{"error": "Like duplicated"})
				return nil
			}
		}

		return err
	}
	ctx.Status(http.StatusOK)
	return nil
}

// LikeDeck godoc
// @Summary      Like a deck
// @Description  Endpoint to like a specific deck. Gives 409 in case of duplicating like
// @Produce      json
// @Param        deckId path string true "Deck ID"
// @Param        userId query string true "User ID"
// @Success      200 {object} map[string]string
// @Failure      400,409 {object} map[string]string
// @Router       /v1/deck/{deckId}/like [post]
// @Router       /v2/deck/{deckId}/like [post]
func (c Controller) LikeDeck(ctx *gin.Context) error {
	deckId := ctx.Param("deckId")
	userId := ctx.Query("userId")

	if _, err := uuid.Parse(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID format for userId",
		})
		return nil
	}

	err := c.services.DecksLikesService.Like(userId, deckId)

	if err != nil {
		var driverErr *mysql.MySQLError
		if errors.As(err, &driverErr) {
			if driverErr.Number == pkg.SqlDuplicateErrState {
				ctx.JSON(http.StatusConflict, gin.H{"error": "Like duplicated"})
				return nil
			}
		}

		return err
	}
	ctx.Status(http.StatusOK)
	return nil
}

// DislikeDeck godoc
// @Summary      DislikeDEPRECATED a deck
// @Description  Endpoint to remove like from a specific deck
// @Produce      json
// @Param        deckId path string true "Deck ID"
// @Param        userId query string true "User ID"
// @Success      200 {object} map[string]string
// @Failure      400,500 {object} map[string]string
// @Router       /v1/deck/{deckId}/dislike [post]
func (c Controller) DislikeDeck(ctx *gin.Context) error {
	deckId := ctx.Param("deckId")
	userId := ctx.Query("userId")

	if _, err := uuid.Parse(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID format for userId",
		})
		return nil
	}

	err := c.services.DecksLikesService.Dislike(userId, deckId)
	if err != nil {
		return err
	}
	ctx.Status(http.StatusOK)
	return nil
}

// GetUserLikesDEPRECATED godoc
// @Summary      Get all likes made by a user
// @Description  Retrieves all likes made by a user on questions and decks.
// @Param 		 userId path string true "The ID of the user."
// @Produce      json
// @Success      200  {object} map[string]interface{}
// @Router       /v1/user/{userId}/likes [get]
func (c Controller) GetUserLikesDEPRECATED(ctx *gin.Context) error {
	userId := ctx.Query("userId")
	if userId == "" {
		userId = ctx.Param("userId")
	}

	if _, err := uuid.Parse(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID format for userId",
		})
		return nil
	}

	dLikes, err := c.services.Repos.DeckLikes.GetAllLikesByUser(userId)
	if err != nil {
		return err
	}

	qLikes, err := c.services.Repos.QuestionReactions.GetAllLikesByUserDEPRECATED(userId)
	if err != nil {
		return err
	}
	answer := make(map[string]any)
	answer["questions"] = qLikes
	answer["decks"] = dLikes

	ctx.JSON(http.StatusOK, answer)
	return nil
}

// GetUserReactions godoc
// @Summary      Get all reactions made by a user
// @Description  Retrieves all reactions made by a user on questions and decks.
// @Param 		 userId path string true "The ID of the user."
// @Produce      json
// @Success      200  {object} map[string]interface{}
// @Router       /v2/user/{userId}/reactions [get]
func (c Controller) GetUserReactions(ctx *gin.Context) error {
	userId := ctx.Query("userId")
	if userId == "" {
		userId = ctx.Param("userId")
	}

	if _, err := uuid.Parse(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID format for userId",
		})
		return nil
	}

	dLikes, err := c.services.Repos.DeckLikes.GetAllLikesByUser(userId)
	if err != nil {
		return err
	}

	qLikes, err := c.services.Repos.QuestionReactions.GetAllReactionsByUser(userId)
	if err != nil {
		return err
	}
	answer := output.NewUserReactions(dLikes, qLikes)

	ctx.JSON(http.StatusOK, answer)
	return nil
}

// EnterPromo godoc
// @Summary      Отправить промокод. Если найдется скрытая колода с таким промо - она будет приходить в /decks
// @Description  Если колода успешно разблокировалась - статус 201 (created) и колода в теле. Если нет промокода - статус 204 (No content) и пустое тело.
// @Param 		 promo path string true "Промокод"
// @Param 		 clientId query string true "Id клиента"
// @Produce      json
// @Success      200  {object} map[string]interface{}
// @Router       /v1/enter-promo/{promo} [post]
func (c Controller) EnterPromo(ctx *gin.Context) error {
	promo := ctx.Param("promo")
	clientID := ctx.Query("clientId")
	if promo == "" {
		return errors.New("promo required")
	}
	if clientID == "" {
		return errors.New("clientId required")
	}
	unlockedDeck, err := c.services.Decks.TryUnlockDeck(promo, clientID)
	if err != nil {
		return err
	}
	if unlockedDeck != nil {
		ctx.JSON(http.StatusCreated, unlockedDeck)
	} else {
		ctx.Status(http.StatusNoContent)
	}
	return nil
}

func doWithErrExplicit(f func(c *gin.Context) error) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := f(context)
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func doWithErr(f func(c *gin.Context) error) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := f(context)
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
	}
}
