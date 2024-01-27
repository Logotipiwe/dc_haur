//go:build wireinject
// +build wireinject

package starter

import (
	"database/sql"
	"dc_haur/src/internal/domain"
	"dc_haur/src/internal/repo"
	"dc_haur/src/internal/service"
	"dc_haur/src/tghttp"
	"github.com/google/wire"
)

func InitServices(db *sql.DB) *service.Services {
	wire.Build(
		repo.NewRepositories,
		tghttp.CreateTgBot,
		domain.NewBotInteractor,
		service.NewServices,
	)
	return &service.Services{}
}
