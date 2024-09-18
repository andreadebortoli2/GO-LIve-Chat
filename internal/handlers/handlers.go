package handlers

import (
	"net/http"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/render"
)

func ShowHomePage(w http.ResponseWriter, r *http.Request) {

	render.RenderPage(w, r, "home", nil)
}
func ShowAboutPage(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"test2": "passing data through handler",
	}

	render.RenderPage(w, r, "about", data)
}
