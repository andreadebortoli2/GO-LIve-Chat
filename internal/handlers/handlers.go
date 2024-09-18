package handlers

import (
	"log"
	"net/http"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	parsePage, err := render.RenderPage()
	if err != nil {
		log.Println(err)
		return
	}

	parsePage.Execute(w, nil)
}
