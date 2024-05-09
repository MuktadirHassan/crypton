package api

import (
	"github.com/MuktadirHassan/crypton/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	additiveHandler := handler.AdditiveCipherHandler{}
	r.Route("/additive", func(r chi.Router) {
		r.Post("/encrypt", additiveHandler.EncryptHandler)
		r.Post("/decrypt", additiveHandler.DecryptHandler)
	})

	vigenereHandler := handler.VigenereCipherHandler{}
	r.Route("/vigenere", func(r chi.Router) {
		r.Post("/encrypt", vigenereHandler.EncryptHandler)
		r.Post("/decrypt", vigenereHandler.DecryptHandler)
	})

	multiplicativeHandler := handler.MultiplicativeCipherHandler{}
	r.Route("/multiplicative", func(r chi.Router) {
		r.Post("/encrypt", multiplicativeHandler.EncryptHandler)
		r.Post("/decrypt", multiplicativeHandler.DecryptHandler)
	})

	affineHandler := handler.AffineCipherHandler{}
	r.Route("/affine", func(r chi.Router) {
		r.Post("/encrypt", affineHandler.EncryptHandler)
		r.Post("/decrypt", affineHandler.DecryptHandler)
	})

	autokeyHandler := handler.AutokeyCipherHandler{}
	r.Route("/autokey", func(r chi.Router) {
		r.Post("/encrypt", autokeyHandler.EncryptHandler)
		r.Post("/decrypt", autokeyHandler.DecryptHandler)
	})

	rsaHandler := handler.RsaCipherHandler{}
	r.Route("/rsa", func(r chi.Router) {
		r.Post("/keygen", rsaHandler.GenerateKeyHandler)
		r.Post("/encrypt", rsaHandler.EncryptHandler)
		r.Post("/decrypt", rsaHandler.DecryptHandler)
	})

	return r
}
