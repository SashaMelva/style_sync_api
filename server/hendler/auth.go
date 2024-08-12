package hendler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/SashaMelva/style_sync_api/internal/auth"
	"github.com/google/uuid"
)

type (
	AuthUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AuthUserIdResponse struct {
		Id uuid.UUID `json:"id"`
	}
)

func (s *Service) Registration(w http.ResponseWriter, r *http.Request) {
	var auth AuthUserRequest
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := json.Unmarshal([]byte(r.Body.Close().Error()), &auth); err != nil {
		ErrUnmarshalJson(err)
		return
	}

	err := s.app.AuthenticationApi.Registration(ctx, auth.schemaAuthUser())

	if err != nil {
		ErrUnmarshalJson(err)
		return
	}

	SucsessRequest("OK", nil)
}

func (s *Service) Authorization(w http.ResponseWriter, r *http.Request) {
	var auth AuthUserRequest
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := json.Unmarshal([]byte(r.Body.Close().Error()), &auth); err != nil {
		ErrUnmarshalJson(err)
		return
	}

	err := s.app.AuthenticationApi.Authorization(ctx, auth.schemaAuthUser())

	if err != nil {
		ErrUnmarshalJson(err)
		return
	}

	SucsessRequest("OK", nil)
}

func (authUser AuthUserRequest) schemaAuthUser() auth.AuthUser {
	return auth.AuthUser{
		Email:    authUser.Email,
		Password: authUser.Password,
	}
}
