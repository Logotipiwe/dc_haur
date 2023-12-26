package main

import (
	"database/sql"
	"dc_haur/src/http"
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/repo"
	"dc_haur/src/internal/service"
	"dc_haur/src/pkg"
	"dc_haur/src/tghttp"
	config "github.com/logotipiwe/dc_go_config_lib"
)

func main() {
	err, db := initializeApp()

	bot := tghttp.CreateTgBot()
	repos := repo.NewRepositories(db)
	services := service.NewServices(repos, domain.NewBotInteractor(bot))

	go http.StartServer(services)

	tghttp.HandleBotUpdates(services)

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
