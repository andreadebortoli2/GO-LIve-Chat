package handlers

import (
	"log"
	"net/http"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/config"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) ShowHomePage(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)

	render.RenderPage(w, r, "home", nil)
}
func (m *Repository) ShowAboutPage(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]string)
	data["test2"] = "passing data through handler"

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	data["remoteIP"] = remoteIP

	render.RenderPage(w, r, "about", data)
}
func (m *Repository) ShowContactPage(w http.ResponseWriter, r *http.Request) {

	render.RenderPage(w, r, "contact", nil)
}
func (m *Repository) PostContact(w http.ResponseWriter, r *http.Request) {
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

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	data["remoteIP"] = remoteIP

	render.RenderPage(w, r, "contact", data)
}
