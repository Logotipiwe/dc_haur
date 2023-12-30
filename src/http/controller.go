package http

import (
	"dc_haur/src/internal/repo"
	"dc_haur/src/internal/service"
	"errors"
	"github.com/Logotipiwe/dc_go_auth_lib/auth"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	config "github.com/logotipiwe/dc_go_config_lib"
	"image/png"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const IntegrationTestPrefix = "/api/v1/integration-test"

func StartServer(services *service.Services) {
	router := gin.Default()
	integrationTestingRoutes := router.Group(IntegrationTestPrefix)

	integrationTestingRoutes.Use(func(c *gin.Context) {
		if err := auth.AuthAsMachine(c.Request); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	})

	integrationTestingRoutes.GET("/test-image", doWithErrExplicit(func(c *gin.Context) error {
		card, err := service.CreateImageCard("Отвечает человек слева: Как ты думаешь, что самое сложное в том деле, которым я зарабатываю себе на жизнь?")
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

	apiV1.GET("/decks", doWithErr(func(c *gin.Context) error {
		decks, err := services.Decks.GetDecks()
		if err != nil {
			return err
		}
		c.JSON(http.StatusOK, decks)
		return nil
	}))

	apiV1.GET("/levels", doWithErr(func(c *gin.Context) error {
		deckId := c.Query("deckId")
		levels, err := services.Questions.GetLevels(deckId)
		if errors.Is(err, repo.NoLevelsErr) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No levels found by deck id " + deckId})
			return nil
		} else if err != nil {
			return err
		}
		c.JSON(http.StatusOK, levels)
		return nil
	}))

	apiV1.GET("/question", doWithErr(func(c *gin.Context) error {
		deckId := c.Query("deckId")
		levelName := c.Query("levelName")
		question, err := services.Questions.GetRandQuestion(deckId, levelName)
		if err != nil {
			return err
		}

		if clientId, exists := c.GetQuery("clientId"); exists {
			err := services.Repos.History.Insert(clientId, question)
			if err != nil {
				return err
			}
		}

		c.JSON(http.StatusOK, question)
		return nil
	}))

	port := config.GetConfigOr("CONTAINER_PORT", "80")
	log.Println("Starting server on port " + port)
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
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
