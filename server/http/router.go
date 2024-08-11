package http

import (
	"fmt"
	"net/http"

	"github.com/SashaMelva/style_sync_api/server/hendler"
	"github.com/go-chi/chi"
)

func weatherReportRouter(h *hendler.Service) chi.Router {
	r := chi.NewRouter()
	r.Get("/by/week", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: index"))
	})
	r.Get("/by/month", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: list accounts.."))
	})
	return r
}

func userRouter(h *hendler.Service) chi.Router {
	r := chi.NewRouter()
	r.Route("/profile", func(r chi.Router) {
		r.Get("/", h.GetProfile)
		r.Put("/", h.UpdateProfile)
	})

	return r
}

func adminRouter(h *hendler.Service) chi.Router {
	r := chi.NewRouter()
	r.Use(AdminOnly)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: index"))
	})
	r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: list accounts.."))
	})
	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("admin: view user id %v", chi.URLParam(r, "userId"))))
	})
	return r
}
