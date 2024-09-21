package handlers

import (
	"log"
	"net/http"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/config"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/database"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/models"
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

// ShowHomePage show home page
func (m *Repository) ShowHomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "home", nil)
}

// ShowAboutPage show about page
func (m *Repository) ShowAboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "about", nil)
}

// ShowLoginPage show login page
func (m *Repository) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, r, "login", nil)
}

// PostLogin logic to login the usesr
func (m *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		// TODO: add message to tell the error to user?
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := database.Login(email, password, r)
	if err != nil {
		log.Println(err)
		// TODO: add message to tell the error to user?
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	if user == (models.User{}) {
		m.App.Session.Put(r.Context(), "auth", nil)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	_ = m.App.Session.RenewToken(r.Context())
	m.App.Session.Put(r.Context(), "auth", user.AccessLevel)

	log.Println("auth:", user.AccessLevel)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ShowLogoutPage logic to logout the user
func (m *Repository) ShowLogoutPage(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.Destroy(r.Context())
	_ = m.App.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
