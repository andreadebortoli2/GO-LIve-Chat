package handlers

import (
	"net/http"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/render"
)

func ShowHomePage(w http.ResponseWriter, r *http.Request) {

	render.RenderPage(w, r, "home")
}
func ShowAboutPage(w http.ResponseWriter, r *http.Request) {

	render.RenderPage(w, r, "about")
}
