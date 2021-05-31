package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"test.com/test/handler"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Attempts to recover from panics with 500 status codes
	r.Use(middleware.Recoverer)
	// CORS
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.RequestID)

	// Bind routes
	r.Get("/twitter/{username}", handler.HandleTwitter)
	r.Get("/tiktok/{username}", handler.HandleTiktok)
	return r
}
