package weatherreport

import (
	"context"
)

type ApiConfig struct {
	WeatherReportReader WeatherReportReader
}

type Api struct {
	config ApiConfig
}

func NewApi(c ApiConfig) (*Api, error) {
	return &Api{config: ApiConfig{
		WeatherReportReader: c.WeatherReportReader,
	}}, nil
}

func (api Api) GetByWeek(ctx context.Context) (*WeatherReport, error) {
	return api.config.WeatherReportReader.GetByWeek(ctx)
}
func (api Api) GetByMounth(ctx context.Context) (*WeatherReport, error) {
	return api.config.WeatherReportReader.GetByMounth(ctx)
}
