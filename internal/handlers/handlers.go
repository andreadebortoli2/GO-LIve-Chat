package handlers

import (
	"log"
	"net/http"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	parsePage, err := render.RenderPage("home")
	if err != nil {
		log.Println(err)
		return
	}

	parsePage.Execute(w, nil)
}
func AboutPage(w http.ResponseWriter, r *http.Request) {

	parsePage, err := render.RenderPage("about")
	if err != nil {
		log.Println(err)
		return
	}

	parsePage.Execute(w, nil)
}
