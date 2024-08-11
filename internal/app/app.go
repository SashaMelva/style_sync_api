package app

import (
	httpweatherreport "github.com/SashaMelva/style_sync_api/adapter/http"
	"github.com/SashaMelva/style_sync_api/adapter/pg"
	"github.com/SashaMelva/style_sync_api/internal/auth"
	"github.com/SashaMelva/style_sync_api/internal/profile"
	weatherreport "github.com/SashaMelva/style_sync_api/internal/weather_report"
	"go.uber.org/zap"
)

type Config struct {
}
type App struct {
	storage             *pg.Storage
	Logger              *zap.SugaredLogger
	AuthenticationApi   *auth.Api
	ProfileUserApi      *profile.Api
	WeatherReportClient *weatherreport.Api
}

func New(logger *zap.SugaredLogger, storage *pg.Storage) *App {
	apiAuth, err := auth.NewApi(
		auth.ApiConfig{
			Authentication: storage,
		})

	if err != nil {
		logger.Fatal(err)
		return nil
	}

	profileAuth, err := profile.NewApi(
		profile.ApiConfig{
			ProfileReader: storage,
			ProfileWriter: storage,
		})

	if err != nil {
		logger.Fatal(err)
		return nil
	}

	client := httpweatherreport.NewClient()

	weatherReportApi, err := weatherreport.NewApi(
		weatherreport.ApiConfig{
			WeatherReportReader: client,
		})

	if err != nil {
		logger.Fatal(err)
		return nil
	}

	return &App{
		storage:             storage,
		Logger:              logger,
		AuthenticationApi:   apiAuth,
		ProfileUserApi:      profileAuth,
		WeatherReportClient: weatherReportApi,
	}
}
