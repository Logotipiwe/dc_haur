package http

import (
	"dc_haur/src/internal/service"
	"image/png"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		err, card := service.CreateImageCard("Отвечает человек слева: Как ты думаешь, что самое сложное в том деле, которым я зарабатываю себе на жизнь?")
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

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}

func sendErr(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
