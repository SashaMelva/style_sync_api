package weatherreport

import (
	"context"
	"time"
)

type WeatherReport struct {
	Report []ReportForDay
}

type ReportForDay struct {
	Date        time.Time
	Temperature int
}

type WeatherReportReader interface {
	GetByWeek(ctx context.Context) (*WeatherReport, error)
	GetByMounth(ctx context.Context) (*WeatherReport, error)
}
