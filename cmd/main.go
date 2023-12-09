package main

import (
	"database/sql"
	"dc_haur/src/pkg"
	"dc_haur/src/tghttp"
	"github.com/logotipiwe/dc_go_utils/src/config"
)

func main() {
	err, db := initializeApp()

	tghttp.StartTgBot(db)

	println("Tg bot started!")
	println("Server up!")
	if err != nil {
		panic("Lol server fell")
	}
}

func initializeApp() (error, *sql.DB) {
	config.LoadDcConfig()
	err, db := pkg.InitDb()
	if err != nil {
		panic(err)
	}
	return err, db
}
