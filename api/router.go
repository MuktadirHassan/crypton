package api

import (
	"github.com/MuktadirHassan/crypton/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/additive", AdditiveCipherRoutes)

	return r
}

func AdditiveCipherRoutes(r chi.Router) {
	additiveCipherHandlers := &handler.AdditiveCipherHandler{}

	r.Post("/encrypt", additiveCipherHandlers.EncryptHandler)
	r.Post("/decrypt", additiveCipherHandlers.DecryptHandler)
}
