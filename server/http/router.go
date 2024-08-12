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

func wardrobeRouter(h *hendler.Service) chi.Router {
	r := chi.NewRouter()
	r.Route("/wardrobe", func(r chi.Router) {
		r.Get("/", h.GetWardrobe)
		r.Get("/category", h.GetWardrobeByCategory)
		r.Post("/", h.AddWardrobe)
		r.Delete("/", h.DeleteWardrobe)
		r.Put("/", h.UpdateWardrobe)

		r.Post("/clothe", h.AddClotheToWardrobe)
		r.Delete("/clothe", h.DeleteClotheToWardrobe)
	})

	return r
}
