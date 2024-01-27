package http

import (
	"dc_haur/src/internal/repo"
	"dc_haur/src/internal/service"
	"errors"
	"github.com/Logotipiwe/dc_go_auth_lib/auth"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	config "github.com/logotipiwe/dc_go_config_lib"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"image/png"
	"log"
	"net/http"

	_ "dc_haur/docs"
	"github.com/gin-gonic/gin"
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
		card, err := service.CreateImageCard("Отвечает человек слева: Как ты думаешь, что самое сложное в том деле, которым я зарабатываю себе на жизнь?", "", "")
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

	apiV1.GET("/decks", doWithErr(controller.GetDecks()))

	apiV1.GET("/levels", doWithErr(controller.GetLevels()))

	apiV1.GET("/question", doWithErr(controller.GetQuestion()))

	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := config.GetConfigOr("CONTAINER_PORT", "80")
	log.Println("Starting server on port " + port)
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

// GetDecks godoc
// @Summary      Get all available decks
// @Produce      json
// @Success      200  {array} domain.Deck
// @Router       /decks [get]
func (c Controller) GetDecks() func(c *gin.Context) error {
	return func(ctx *gin.Context) error {

		decks, err := c.services.Decks.GetDecks()
		if err != nil {
			return err
		}
		ctx.JSON(http.StatusOK, decks)
		return nil
	}
}

// GetLevels godoc
// @Summary      Get levels from specified deck
// @Param 		 deckId query string true "Id of deck for which selecting levels"
// @Produce      json
// @Success      200  {array} domain.Level
// @Router       /levels [get]
func (c Controller) GetLevels() func(ctx *gin.Context) error {
	return func(ctx *gin.Context) error {
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
}

// GetQuestion godoc
// @Summary      Get random question by selected level
// @Param 		 levelId query string true "Id of level for which selecting question"
// @Param 		 clientId query string true "Client id - differs clients from each other. Needed for ordering random questions for each client/"
// @Produce      json
// @Success      200  {object} domain.Question
// @Router       /question [get]
func (c Controller) GetQuestion() func(ctx *gin.Context) error {
	return func(ctx *gin.Context) error {
		levelID := ctx.Query("levelId")
		question, err := c.services.Questions.GetRandQuestion(levelID)
		if err != nil {
			return err
		}

		if clientId, exists := ctx.GetQuery("clientId"); exists {
			err := c.services.Repos.History.Insert(clientId, question)
			if err != nil {
				return err
			}
		}

		ctx.JSON(http.StatusOK, question)
		return nil
	}
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
