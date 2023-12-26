package http

import (
	"dc_haur/src/internal/service"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	config "github.com/logotipiwe/dc_go_config_lib"
	"image/png"
	"log"
	"net/http"
)

func StartServer(services *service.Services) {
	http.HandleFunc("/test-image", func(w http.ResponseWriter, r *http.Request) {
		card, err := service.CreateImageCard("Отвечает человек слева: Как ты думаешь, что самое сложное в том деле, которым я зарабатываю себе на жизнь?")
		if err != nil {
			sendErr(w, err)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		err = png.Encode(w, card)
		if err != nil {
			sendErr(w, err)
			return
		}
	})

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		var update tgbotapi.Update
		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			sendErr(w, err)
			return
		}
		reply, err := services.TgUpdatesHandler.HandleMessageAndReply(update)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			reply = services.TgUpdatesHandler.SendUnknownCommandAnswer(update)
		}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/clear-history", func(w http.ResponseWriter, r *http.Request) {
		err := services.Repos.History.Truncate()
		if err != nil {
			sendErr(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/images-enabled", func(w http.ResponseWriter, r *http.Request) {
		enabledImagesStr := config.GetConfig("ENABLE_IMAGES")
		_, err := fmt.Fprint(w, enabledImagesStr)
		if err != nil {
			log.Print(err)
		}
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}

func sendErr(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
