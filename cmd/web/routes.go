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
	// load session
	r.Use(LoadSession)

	// public routes
	r.Get("/", handlers.Repo.ShowHomePage)
	r.Get("/about", handlers.Repo.ShowAboutPage)
	r.Get("/login", handlers.Repo.ShowLoginPage)
	r.Post("/login", handlers.Repo.PostLogin)
	r.Get("/logout", handlers.Repo.ShowLogoutPage)
	r.Get("/new-user", handlers.Repo.ShowNewUserPage)
	r.Post("/new-user", handlers.Repo.PostNewUserPage)

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

// LoadSession load the session
func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
