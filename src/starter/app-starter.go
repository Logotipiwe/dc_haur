package starter

import (
	"database/sql"
	"dc_haur/docs"
	"dc_haur/src/http"
	"dc_haur/src/pkg"
	"dc_haur/src/tghttp"
	config "github.com/logotipiwe/dc_go_config_lib"
	"os"
)

func StartApp() {
	err, db := initializeApp()

	setSwaggerHost()

	services := InitServices(db)

	go http.StartServer(services)

	tghttp.HandleBotUpdates(services)

	println("Tg bot started!")
	println("Server up!")
	if err != nil {
		panic("Lol server fell")
	}
}

func setSwaggerHost() {
	ns := os.Getenv("NAMESPACE")
	if ns == "" {
		//prod
		docs.SwaggerInfo.Host = "logotipiwe.ru"
		docs.SwaggerInfo.BasePath = "/haur/api"
	} else {
		docs.SwaggerInfo.Host = "localhost:" + config.GetConfigOr("CONTAINER_PORT", "80")
		docs.SwaggerInfo.BasePath = "/api"
	}
}

func initializeApp() (error, *sql.DB) {
	config.LoadDcConfigDynamically(3)
	err, db := pkg.InitDb()
	if err != nil {
		panic(err)
	}
	return err, db
}
