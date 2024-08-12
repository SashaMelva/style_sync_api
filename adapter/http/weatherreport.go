package httpweatherreport

import (
	"context"
	"net/http"

	weatherreport "github.com/SashaMelva/style_sync_api/internal/weather_report"
)

type Client struct {
	client http.Client
}

func NewClient() *Client {
	client := http.Client{}

	return &Client{
		client: client,
	}
}

func (c Client) GetByWeek(ctx context.Context) (*weatherreport.WeatherReport, error) {
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// var schema WeatherReportSchema

	// res, err := json.Unmarshal(resp, &schema)

	// if err != nil {
	// 	return nil, err
	// }

	// return res, nil
	return nil, nil
}
func (c Client) GetByMounth(ctx context.Context) (*weatherreport.WeatherReport, error) {
	return nil, nil
}

type WeatherReportSchema struct {
}
