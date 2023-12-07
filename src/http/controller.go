package http

import (
	"database/sql"
	"dc_haur/src/internal/service"
	"fmt"
	"log"
	"net/http"
)

func StartHttpServer(*sql.DB) {

	log.Print("SERVER STARTING...")

	generator := service.ImageGenerator{}
	http.HandleFunc("/gradient", generator.HandleGradientRequest)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
