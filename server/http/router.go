package http

import (
	"github.com/SashaMelva/style_sync_api/server/hendler"
	"github.com/go-chi/chi"
)

func weatherReportRouter(h *hendler.Service) chi.Router {
	r := chi.NewRouter()
	r.Get("/by/week", h.GetByWeek)
	r.Get("/by/month", h.GetByMounth)
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
