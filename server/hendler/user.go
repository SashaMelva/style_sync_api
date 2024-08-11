package hendler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/SashaMelva/style_sync_api/server/schema"
	"github.com/go-chi/render"
)

func (s *Service) Registration(w http.ResponseWriter, r *http.Request) {
	data := &schema.NewUserRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewArticleResponse(article))
}

func (user *schema.NewUserRequest) Bind(r *http.Request) error {
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return errors.New("missing required Article fields.")
	}
	if user.Password == "" {
		return errors.New("missing required Article fields.")
	}

	return nil
}

func (s *Service) Authorization(w http.ResponseWriter, r *http.Request) {

}

func (s *Service) GetProfile(w http.ResponseWriter, r *http.Request) {

}

func (s *Service) UpdateProfile(w http.ResponseWriter, r *http.Request) {

}
