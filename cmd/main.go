package main

import (
	"database/sql"
	"dc_haur/src/http"
	"dc_haur/src/pkg"
	"dc_haur/src/tghttp"
	config "github.com/logotipiwe/dc_go_config_lib"
)

func main() {
	err, db := initializeApp()

	go http.StartServer()

	tghttp.StartTgBot(db)

	println("Tg bot started!")
	println("Server up!")
	if err != nil {
		panic("Lol server fell")
	}
}

func initializeApp() (error, *sql.DB) {
	config.LoadDcConfigDynamically(3)
	err, db := utils.InitDb()
	if err != nil {
		panic(err)
	}
	return err, db
}
