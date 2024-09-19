package handlers

import (
	"log"
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
func ShowContactPage(w http.ResponseWriter, r *http.Request) {

	render.RenderPage(w, r, "contact", nil)
}
func PostContact(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		log.Println("not working")
		return
	}
	log.Println("working")

	data := map[string]string{
		"success": "form posted successfully",
	}
	render.RenderPage(w, r, "contact", data)
}
