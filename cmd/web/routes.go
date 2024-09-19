package main

import (
	"net/http"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/justinas/nosurf"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	// MIDDLEWARES
	// recover from panics, log the panic and return an HTTP 500 status(if possible)
	r.Use(middleware.Recoverer)
	// add crsf token in cookies
	r.Use(NoSurf)
	// TODO! load session

	r.Get("/", handlers.ShowHomePage)
	r.Get("/about", handlers.ShowAboutPage)
	r.Get("/contact", handlers.ShowContactPage)
	r.Post("/contact", handlers.PostContact)

	return r
}

// MIDDLEWARES
// NoSurf generate CRSF token
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
