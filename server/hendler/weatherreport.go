package hendler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (s *Service) GetByWeek(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	report, err := s.app.WeatherReportClient.GetByWeek(ctx)

	if err != nil {
		ErrInternalServer(err)
		return
	}

	res, err := json.Marshal(report)
	if err != nil {
		ErrUnmarshalJson(err)
		return
	}

	SucsessRequest("OK", &res)
}

func (s *Service) GetByMounth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	report, err := s.app.WeatherReportClient.GetByMounth(ctx)

	if err != nil {
		ErrInternalServer(err)
		return
	}

	res, err := json.Marshal(report)
	if err != nil {
		ErrUnmarshalJson(err)
		return
	}

	SucsessRequest("OK", &res)
}
