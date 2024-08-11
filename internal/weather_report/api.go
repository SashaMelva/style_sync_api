package weatherreport

import (
	"context"
)

type ApiConfig struct {
	WeatherReportReader WeatherReportReader
}

type Api struct {
	weatherReportReader WeatherReportReader
}

func NewApi(c ApiConfig) (*Api, error) {
	return &Api{
		weatherReportReader: c.WeatherReportReader,
	}, nil
}

func (api Api) GetByWeek(ctx context.Context) (*WeatherReport, error)   { return nil, nil }
func (api Api) GetByMounth(ctx context.Context) (*WeatherReport, error) { return nil, nil }
