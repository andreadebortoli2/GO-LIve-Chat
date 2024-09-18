package main

import (
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	// midlewares
	// r.Use()

	r.Get("/", handlers.ShowHomePage)
	r.Get("/about", handlers.ShowAboutPage)

	return r
}
