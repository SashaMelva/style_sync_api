package hendler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/SashaMelva/style_sync_api/internal/profile"
	"github.com/google/uuid"
)

type profileSchema struct {
	Id uuid.UUID
}

func (s *Service) GetProfile(w http.ResponseWriter, r *http.Request) {
	userId := uuid.New()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	profile, err := s.app.ProfileUserApi.GetProfile(ctx, userId)

	if err != nil {
		ErrInternalServer(err)
		return
	}

	res, err := json.Marshal(profile)
	if err != nil {
		ErrUnmarshalJson(err)
		return
	}

	SucsessRequest("OK", &res)
}

func (s *Service) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var schema profileSchema

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := json.Unmarshal([]byte(r.Body.Close().Error()), &schema); err != nil {
		ErrUnmarshalJson(err)
		return
	}

	profile := schema.schemaProfile()
	err := s.app.ProfileUserApi.UpdateProfile(ctx, profile)

	if err != nil {
		ErrInternalServer(err)
		return
	}

	res, err := json.Marshal(profile)
	if err != nil {
		ErrUnmarshalJson(err)
		return
	}

	SucsessRequest("OK", &res)
}

func (schema profileSchema) schemaProfile() profile.Profile {
	return profile.Profile{
		Id: schema.Id,
	}
}
