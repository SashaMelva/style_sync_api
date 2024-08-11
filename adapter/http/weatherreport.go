package httpweatherreport

import (
	"context"

	weatherreport "github.com/SashaMelva/style_sync_api/internal/weather_report"
)

type Client struct {
}

func NewClient() *Client {
	return nil
}

func (c Client) GetByWeek(ctx context.Context) (*weatherreport.WeatherReport, error) {
	return nil, nil
}
func (c Client) GetByMounth(ctx context.Context) (*weatherreport.WeatherReport, error) {
	return nil, nil
}
