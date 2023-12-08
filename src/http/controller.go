package http

import (
	"database/sql"
	"dc_haur/src/internal/service"
	"fmt"
	"image/png"
	"log"
	"net/http"
)

func StartHttpServer(*sql.DB) {

	log.Print("SERVER STARTING...")

	generator := service.ImageGenerator{}
	http.HandleFunc("/gradient", func(w http.ResponseWriter, r *http.Request) {
		image := generator.HandleGradientRequest(w, r)
		// Set the Content-Type header
		w.Header().Set("Content-Type", "image/png")

		// Encode the image and write it to the response writer
		err := png.Encode(w, image)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
